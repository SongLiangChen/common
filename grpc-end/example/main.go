package main

import (
	"github.com/becent/golang-common/grpc-end"
	"github.com/becent/golang-common/grpc-end/middleware"
	"google.golang.org/grpc"
	"math"
	"time"
)

func main() {
	ser, err := newGRpcEngine().Run(":9090", grpc.MaxRecvMsgSize(math.MaxInt32))
	if err != nil {
		println(err.Error())
		return
	}

	// hold here and deal request...
	time.Sleep(time.Second * 100)

	ser.GracefulStop()
}

func newGRpcEngine() *grpc_end.GRpcEngine {
	engine := grpc_end.NewGRpcEngine("MyAppName")
	engine.RegisterFunc("hello", "world", sayHi)

	engine.Use(middleware.Recover)
	engine.Use(middleware.Logger)

	return engine
}

func sayHi(c *grpc_end.GRpcContext) {
	name := c.StringParamDefault("name", "Tom")
	c.SuccessResponse("Hi" + name)
}
