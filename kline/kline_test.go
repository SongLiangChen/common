package kline

import (
	"strconv"
	"testing"
	"time"
)

func TestEngine_Run(t *testing.T) {
	engine := NewEngine()
	engine.SetCacheLimit(1000)

	go func() {
		t := time.NewTicker(time.Second)
		defer t.Stop()
		for {
			select {
			case <-t.C:
				now := strconv.FormatInt(time.Now().Unix(), 10)
				k := &Kline{
					CoinType:   "USDT/BTC",
					High:       now,
					Low:        now,
					Open:       now,
					Close:      now,
					CreateTime: time.Now().Unix(),
					Volume:     10,
				}

				engine.GetSimplingChan() <- k

			case <-engine.exit:
				return
			}
		}

	}()

	engine.SetScale([]int64{1, 60, 300})
	engine.SetHandler(func(kline *Kline) error {
		// println("save to db")
		return nil
	})
	engine.SetHandler(func(kline *Kline) error {
		// println("push to mq")
		return nil
	})
	engine.SetHandler(func(kline *Kline) error {
		// println("increment update data:")
		// incUpData := engine.IncrementalUpdateKline(kline)
		// for _, k := range incUpData {
		// println(k.TimeScale, k.Open)
		// }
		return nil
	})
	if err := engine.Run(); err != nil {
		t.Fatal(err.Error())
	}

	exit := make(chan bool)

	go func() {
		ticker := time.NewTicker(time.Second * 5)
		defer ticker.Stop()

		for {
			select {
			case <-exit:
				return

			case <-ticker.C:
				n, _ := engine.GetKlineList("USDT/BTC", 1, time.Now().Unix(), 100)
				println(n)
			}
		}
	}()

	time.Sleep(time.Minute)
	close(exit)
	engine.Stop()
}
