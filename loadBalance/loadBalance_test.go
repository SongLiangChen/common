package loadBalance

import (
	"github.com/SongLiangChen/common/loadBalance/roundRobin"
	"github.com/SongLiangChen/common/registry"
	"github.com/SongLiangChen/common/registry/etcd"
	"strconv"
	"testing"
	"time"
)

func TestRemainder(t *testing.T) {
	var reg registry.Registry
	reg = etcd.NewRegistry(registry.Timeout(time.Second*3), registry.Addrs("127.0.0.1:2379"))
	go func() {
		for i := 1; i <= 9; i++ {
			ser := &registry.Service{
				Name:  "test",
				Nodes: []*registry.Node{&registry.Node{Id: strconv.Itoa(i), Address: "127.0.0.1:8" + strconv.Itoa(i)}},
			}
			if err := reg.Register(ser, registry.RegisterTTL(time.Second*10)); err != nil {
				t.Logf(err.Error())
			}
			time.Sleep(time.Second * 10)
		}
	}()

	// lb := &remainder.RemainderLoadBalance{}
	lb := &roundRobin.RoundRobinLoadBalance{}
	lb.SetServiceName("test")
	lb.SetRegistry(reg)
	lb.SetReloadFunc(func() error {
		println("触发reload")
		return nil
	})

	time.Sleep(time.Second)

	if err := lb.Start(time.Second * 10); err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 100; i++ {
		ser := lb.GetService(strconv.Itoa(i))
		if ser != nil {
			println("get server:", i, ser.Id, ser.Address)
		}
		time.Sleep(time.Second)
	}
}
