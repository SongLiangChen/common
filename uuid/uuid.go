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
func InitUUID() {
	sf = snowflake.NewSnowflake(snowflake.Settings{
		StartTime: time.Time{},
	})
}

// UUID returns a uint64 union id
// error not nil if any incorrect happen
func UUID() (uint64, error) {
	return sf.NextID()
}
