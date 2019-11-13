package main

import (
	"fmt"
	"github.com/becent/golang-common/session"
	_ "github.com/becent/golang-common/session/redis"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.Any("/test", func(c *gin.Context) {
		m := session.GetSessionManager("abc")
		sess, err := m.SessionStart(c.Writer, c.Request)
		defer sess.SessionRelease(c.Writer)

		if err != nil {
			c.Data(200, "text", []byte(err.Error()))
			return
		}

		if name := sess.Get("name"); name != nil {
			c.Data(200, "text", []byte(name.(string)))
			return
		}

		sess.Set("name", "bob")
		c.Data(200, "text", []byte("set name bob"))
	})

	pcnf := fmt.Sprintf("%s:%s,%d,%s,%d,180", "127.0.0.1", "6379", 100, "", 1)
	println(pcnf)
	err := session.RegisterSessionManager("abc", "redis", &session.ManagerConfig{
		CookieName:      "sessionId",
		EnableSetCookie: true,
		Maxlifetime:     3600,
		Secure:          false,
		ProviderConfig:  pcnf,
		SessionIDPrefix: "abc_",
		CookieLifeTime:  0,
	})
	if err != nil {
		println(err.Error())
		return
	}

	engine.Run(":8080")

}
