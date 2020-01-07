package loadBalance

import (
	"github.com/SongLiangChen/common/registry"
	"time"
)

type LoadBalance interface {
	Ready() bool
	SetServiceName(name string)
	SetRegistry(reg registry.Registry)
	SetReloadFunc(func() error)
	GetService(key string) *registry.Node
	Start(TTL time.Duration) error
}
