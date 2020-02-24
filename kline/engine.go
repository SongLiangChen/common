package kline

import (
	"errors"
	"fmt"
	"github.com/SongLiangChen/common"
	"sort"
	"sync"
)

type Engine struct {
	// 初始化函数，用来做kline恢复啥的特别好使
	initFunc func(e *Engine) error

	scale []int64 // 配置的刻度,按秒为刻度

	simplingChan chan *Kline // 数据采样管道

	// 数据处理链，当采样成功后，数据将当作参数按顺序进行调用
	// 发生任何错误将会中断
	klineHandlers []func(kline *Kline) error

	// 数据适配函数，设置该函数后，当请求kline列表时，会按照该函数进行转化
	klineAdapter func(kline *Kline) interface{}

	// 数据缓存，key是coinType_scale组合
	cachedKline map[string][]*Kline
	// 缓存数量
	cacheLimit int

	exit chan bool

	sync.RWMutex
	common.WaitGroupWrapper
}

func NewEngine() *Engine {
	return &Engine{
		scale:         make([]int64, 0),
		simplingChan:  make(chan *Kline, 1024),
		klineHandlers: make([]func(kline *Kline) error, 0),
		cachedKline:   make(map[string][]*Kline),
		exit:          make(chan bool),
	}
}

// 设置初始化函数，在调用engine.Run时第一个被调用
func (e *Engine) SetInitFunc(f func(e *Engine) error) {
	e.initFunc = f
}

// 设置kline刻度，调用Run方法前必须设置
func (e *Engine) SetScale(scale []int64) {
	e.scale = scale
}

// 获取采样管道，获取到k线数据后，直接写入该管道即可
func (e *Engine) GetSimplingChan() chan<- *Kline {
	return e.simplingChan
}

// 设置新采样k线数据的handler
func (e *Engine) SetHandler(f func(kline *Kline) error) {
	e.klineHandlers = append(e.klineHandlers, f)
}

// 设置k线适配函数，如果不设置，将按默认kline格式输出
func (e *Engine) SetAdapter(f func(kline *Kline) interface{}) {
	e.klineAdapter = f
}

// 获取k线在各个刻度的增量更新
func (e *Engine) IncrementalUpdateKline(k *Kline) []*Kline {
	sks := e.klineForAllScale(k)

	e.RLock()
	defer e.RUnlock()

	ret := make([]*Kline, 0)

	for _, sk := range sks {
		var (
			key  = fmt.Sprintf("%v_%v", sk.CoinType, sk.TimeScale)
			n    = len(e.cachedKline[key])
			last *Kline
		)

		if n > 0 {
			last = e.cachedKline[key][n-1].Copy()
		}

		if last == nil {
			ret = append(ret, sk.Copy())
			continue
		}

		// 更新最后一条数据
		if last.CreateTime == sk.CreateTime {
			if ret, _ := common.BcCmp(last.High, sk.High); ret < 0 {
				last.High = sk.High
			}
			if ret, _ := common.BcCmp(last.Low, sk.Low); ret > 0 {
				last.Low = sk.Low
			}
			last.Close = sk.Close
			last.Volume += sk.Volume // 累加24小时交易量
			ret = append(ret, last)
			continue
		}

		// 同一个24小时
		if last.CreateTime/24*60*60 == sk.CreateTime/24*60*60 {
			sk.Volume += last.Volume // 累加24小时交易量
		}
		ret = append(ret, sk.Copy())

	}

	return ret
}

// 设置缓存限值
func (e *Engine) SetCacheLimit(l int) {
	e.cacheLimit = l
}

// 缓存新的增量更新数据
func (e *Engine) cacheKlines(ks []*Kline) {
	e.Lock()
	e.Unlock()

	for _, k := range ks {
		key := fmt.Sprintf("%v_%v", k.CoinType, k.TimeScale)
		if _, ok := e.cachedKline[key]; !ok {
			e.cachedKline[key] = make([]*Kline, 0)
		}

		var (
			n    = len(e.cachedKline[key])
			last *Kline
		)

		if n > 0 {
			last = e.cachedKline[key][n-1]
		}

		if last == nil || last.CreateTime != k.CreateTime {
			e.cachedKline[key] = append(e.cachedKline[key], k)
		} else {
			e.cachedKline[key][n-1] = k
		}

		if n > e.cacheLimit {
			e.cachedKline[key] = e.cachedKline[key][n-e.cacheLimit:]
		}
		println(key, len(e.cachedKline[key]))
	}
}

// 按所有时间刻度复制k线数据，并修改k线的时间刻度
func (e *Engine) klineForAllScale(k *Kline) []*Kline {
	ret := make([]*Kline, 0)
	for _, s := range e.scale {
		tmp := k.Copy()
		tmp.CreateTime = tmp.CreateTime - tmp.CreateTime%s
		tmp.TimeScale = s
		ret = append(ret, tmp)
	}
	return ret
}

// 启动引擎
func (e *Engine) Run() error {
	if e.initFunc != nil {
		if err := e.initFunc(e); err != nil {
			return err
		}
	}

	if len(e.scale) == 0 {
		return errors.New("engine's scale not set")
	}
	if e.cacheLimit == 0 {
		e.cacheLimit = 20000
	}

	e.Wrap(e.workLoop)

	return nil
}

// 停止引擎
func (e *Engine) Stop() {
	close(e.exit)
	e.Wait()
}

// 主要工作loop
func (e *Engine) workLoop() {
	for {
		select {
		case k := <-e.simplingChan:
			if k == nil {
				break
			}

			for _, h := range e.klineHandlers {
				if err := h(k); err != nil {
					goto NEXT
				}
			}

			e.cacheKlines(e.IncrementalUpdateKline(k))

		case <-e.exit:
			return
		}

	NEXT:
	}
}

func (e *Engine) GetKlineList(coinType string, timeScale int64, endTime int64, size int) (int, interface{}) {
	key := fmt.Sprintf("%v_%v", coinType, timeScale)

	tmpKlines := make([]*Kline, 0)

	e.RLock()
	klines := e.cachedKline[key]
	if len(klines) == 0 {
		e.RUnlock()
		return 0, []struct{}{}
	}

	// 二分查找
	index := sort.Search(len(klines), func(i int) bool {
		return klines[i].CreateTime >= endTime
	})
	l := index - size
	if l < 0 {
		l = 0
	}
	for i := l; i < index; i++ {
		tmpKlines = append(tmpKlines, klines[i].Copy())
	}

	e.RUnlock()

	if e.klineAdapter == nil {
		return len(tmpKlines), tmpKlines
	}

	ret := make([]interface{}, 0)
	for _, tk := range tmpKlines {
		ret = append(ret, e.klineAdapter(tk))
	}
	return len(ret), ret
}
