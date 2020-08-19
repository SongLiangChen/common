package middleware

import (
	"runtime"

	grpc_end "github.com/SongLiangChen/common/grpc-end"
	log "github.com/sirupsen/logrus"
)

// Recover the first middleware to handle request,
// it catch panic exception of this request, make sure the system healthy and strong,
// and panic exception will be logged to log file.
//
// defer() and recover() will be take about 20ns loss every request, but it still necessary.
func Recover(c *grpc_end.GRpcContext) {
	defer func() {
		if err := recover(); err != nil {
			log.WithFields(log.Fields{
				"app":   c.GetAppName(),
				"stack": stack(),
			}).Error(err)
		}
	}()

	c.Next()
}

func stack() string {
	var buf [2 << 10]byte
	return string(buf[:runtime.Stack(buf[:], true)])
}

func RecoverWithCb(cb ...func(c *grpc_end.GRpcContext, stack string)) func(c *grpc_end.GRpcContext) {
	return func(c *grpc_end.GRpcContext) {
		defer func() {
			if err := recover(); err != nil {
				s := stack()

				log.WithFields(log.Fields{
					"app":   c.GetAppName(),
					"stack": s,
				}).Error(err)

				for _, f := range cb {
					f(c, s)
				}
			}
		}()

		c.Next()
	}
}
