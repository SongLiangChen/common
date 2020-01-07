// package roundRobin provides a round robin balance loader
package roundRobin

import (
	"github.com/SongLiangChen/common"
	"github.com/SongLiangChen/common/registry"
	"sort"
	"sync"
	"sync/atomic"
	"time"
)

// 轮训模式负载均衡实现
type RoundRobinLoadBalance struct {
	ServiceName string
	Reg         registry.Registry
	ReadyState  int32
	ReloadFunc  func() error
	Nodes       []*registry.Node

	index int32

	sync.RWMutex
}

func (b *RoundRobinLoadBalance) SetServiceName(name string) {
	b.ServiceName = name
}

func (b *RoundRobinLoadBalance) SetRegistry(reg registry.Registry) {
	b.Reg = reg
}

func (b *RoundRobinLoadBalance) SetReloadFunc(f func() error) {
	b.ReloadFunc = f
}

func (b *RoundRobinLoadBalance) Ready() bool {
	return atomic.LoadInt32(&b.ReadyState) == 1
}

func (b *RoundRobinLoadBalance) GetService(key string) *registry.Node {
	i := atomic.AddInt32(&b.index, 1)

	b.RLock()
	defer b.RUnlock()

	i = i % int32(len(b.Nodes))
	return b.Nodes[int(i)]
}

func (b *RoundRobinLoadBalance) Start(TTL time.Duration) error {
	if b.ServiceName == "" {
		panic("serviceName empty")
	}
	if b.Reg == nil {
		panic("registry is nil")
	}

	if err := b.reload(); err != nil {
		return err
	}

	go b.watch()
	if TTL > 0 {
		go b.keepAlive(TTL)
	}

	return nil
}

func (b *RoundRobinLoadBalance) reload() error {
	atomic.StoreInt32(&b.ReadyState, 0)
	ss, err := b.Reg.GetService(b.ServiceName)
	if err != nil {
		return err
	}

	b.Lock()
	b.Nodes = make([]*registry.Node, 0)
	for _, s := range ss {
		b.Nodes = append(b.Nodes, s.Nodes...)
	}
	sort.Slice(b.Nodes, func(i, j int) bool {
		return b.Nodes[i].Id < b.Nodes[j].Id
	})
	b.Unlock()

	if b.ReloadFunc != nil {
		if err = b.ReloadFunc(); err != nil {
			return err
		}
	}

	atomic.StoreInt32(&b.ReadyState, 1)
	return nil
}

func (b *RoundRobinLoadBalance) watch() {
	var (
		watch     registry.Watcher
		err       error
		noWatcher = true
	)

	for {
		if noWatcher {
			watch, err = b.Reg.Watch(registry.WatchService(b.ServiceName))
			if err == nil {
				noWatcher = false
			} else {
				common.ErrorLog("new watcher err", nil, err.Error())
				continue
			}
		}

		_, err := watch.Next()
		if err != nil {
			common.ErrorLog("load balance watch err", nil, err.Error())
			continue
		}

		if err = b.reload(); err != nil {
			common.ErrorLog("load balance relaod err", nil, err.Error())
		}
	}
}

func (b *RoundRobinLoadBalance) keepAlive(TTL time.Duration) {
	TTL = TTL / 3
	ticker := time.NewTicker(TTL)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			b.Reg.KeepAliveOnce()
		}
	}
}
