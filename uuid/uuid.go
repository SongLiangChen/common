package uuid

import (
	"github.com/becent/golang-common/uuid/snowflake"
	"time"
)

var (
	sf *snowflake.Snowflake
)

// InitUUID init uuid seed
// startTime more closer the time now, the smaller the generated id
// if no startTime set, 2014-09-01 default
func InitUUID(startTime ...time.Time) {
	t := time.Time{}
	if len(startTime) > 0 {
		t = startTime[0]
	}
	sf = snowflake.NewSnowflake(snowflake.Settings{
		StartTime: t,
	})
}

// UUID returns a uint64 union id
// error not nil if any incorrect happen
func UUID() (uint64, error) {
	return sf.NextID()
}
