package remainder

import (
	"github.com/SongLiangChen/common/loadBalance/roundRobin"
	"github.com/SongLiangChen/common/registry"
	"github.com/mitchellh/hashstructure"
)

// 取余模式负载均衡实现
type RemainderLoadBalance struct {
	roundRobin.RoundRobinLoadBalance
}

func (b *RemainderLoadBalance) GetService(key string) *registry.Node {
	id, err := hashstructure.Hash(key, nil)
	if err != nil {
		println(err.Error())
		return nil
	}

	b.RLock()
	defer b.RUnlock()

	return b.Nodes[int(id%uint64(len(b.Nodes)))]
}
