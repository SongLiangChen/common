package common

import (
	"errors"
	"github.com/go-redis/redis"
	"io"
)

type MyRedis interface {
	redis.Cmdable
	io.Closer
}

type MapRedisCache struct {
	rediss map[string]MyRedis
}

func NewMapRedisCache() *MapRedisCache {
	return &MapRedisCache{
		rediss: make(map[string]MyRedis),
	}
}

func (c MapRedisCache) Close() {
	for _, r := range c.rediss {
		r.Close()
	}
}

// AddRedisInstance add a redis instance into poll
// note that, all instance should be added in init status of Application
// dbNum work just if len(addrs)==1
func (c MapRedisCache) AddRedisInstance(key string, addrs []string, pwd string, dbNum int) error {
	if len(addrs) == 0 {
		return errors.New("addrs empty")
	}

	if _, ok := c.rediss[key]; !ok {
		var r MyRedis
		if len(addrs) == 1 {
			r = redis.NewClient(&redis.Options{
				Addr:       addrs[0],
				Password:   pwd,
				DB:         dbNum,
				MaxRetries: 2, // retry 3 times (<=MaxRetries)
				PoolSize:   1024,
			})
		} else {
			r = redis.NewClusterClient(&redis.ClusterOptions{
				Addrs:      addrs,
				Password:   pwd,
				MaxRetries: 2, // retry 3 times (<=MaxRetries)
				PoolSize:   1024,
			})
		}

		if _, err := r.Ping().Result(); err == nil {
			c.rediss[key] = r
		} else {
			return err
		}

	} else {
		return errors.New("repeated key")
	}

	return nil
}

func (c MapRedisCache) GetRedisInstance(key string) (MyRedis, bool) {
	r, ok := c.rediss[key]
	return r, ok
}

// -----------------------------------------------------------------------------

var (
	defaultMapCache *MapRedisCache
)

func init() {
	defaultMapCache = NewMapRedisCache()
}

func AddRedisInstance(key string, addrs []string, pwd string, dbNum int) error {
	return defaultMapCache.AddRedisInstance(key, addrs, pwd, dbNum)
}

func GetRedisInstance(key string) (MyRedis, bool) {
	return defaultMapCache.GetRedisInstance(key)
}

func MustGetRedisInstance(instance ...string) MyRedis {
	name := ""
	if len(instance) == 1 {
		name = instance[0]
	}
	redisInstance, ok := GetRedisInstance(name)
	if !ok {
		panic("redis instance [" + name + "] not exists")
	}

	return redisInstance
}

func ReleaseRedisPool() {
	defaultMapCache.Close()
}
