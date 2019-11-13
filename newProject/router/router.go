package router

import (
	"os"
	"strings"
)

func G_router(projectName string) error {
	if err := os.Mkdir(projectName+"/router", 755); err != nil {
		return err
	}

	file, err := os.OpenFile(projectName+"/router/router.go", os.O_CREATE|os.O_RDWR, 755)
	if err != nil {
		return err
	}

	if _, err := file.WriteString(strings.Replace(router_temple, "{{projectName}}", projectName, -1)); err != nil {
		return err
	}

	return nil

}

var router_temple = `package router

import (
	"{{projectName}}/handler"
	"{{projectName}}/gRpcHandler"
	"{{projectName}}/config"
	"github.com/gin-gonic/gin"
	"github.com/becent/golang-common/gin-handler"
	"github.com/becent/golang-common/grpc-end"
	"github.com/becent/golang-common/grpc-end/middleware"
)

func NewGinEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	engine.Use(gin_handler.NewHandler(&gin_handler.Config{
		AppName:        config.GetConfig("system", "app_name"),
		CheckSignature: false,
	}))
	engine.Any("/", lvsHealthCheck)

	// TODO add your router here...
	engine.GET("/hello", handler.Hello)

	return engine
}

func lvsHealthCheck(c *gin.Context) {
	c.String(200, "%s", "ok")
}

func NewGRpcEngine() *grpc_end.GRpcEngine {
	engine := grpc_end.NewGRpcEngine(config.GetConfig("system", "app_name"))
	engine.RegisterFunc("", "hello", gRpcHandler.Hello)

	// TODO use your middleware here
	engine.Use(middleware.Recover)
	engine.Use(middleware.Logger)

	return engine
}
`
