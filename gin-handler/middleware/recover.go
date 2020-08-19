package middleware

import (
	"runtime"

	gin_handler "github.com/SongLiangChen/common/gin-handler"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// Recover the first gin'middleware to handle request,
// it catch panic exception of this request, make sure the system healthy and strong,
// and panic exception will be logged to log file.
//
// defer() and recover() will be take about 20ns loss every request, but it still necessary.
func Recover(c *gin.Context, pf ...func(*gin.Context)) {
	defer func() {
		if err := recover(); err != nil {
			log.WithFields(log.Fields{
				"app":   c.GetString(gin_handler.KEY_APPNAME),
				"stack": stack(),
			}).Error(err)

			for _, f := range pf {
				f(c)
			}
		}
	}()

	c.Next()
}

func stack() string {
	var buf [2 << 10]byte
	return string(buf[:runtime.Stack(buf[:], true)])
}
