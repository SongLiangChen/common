package main

import (
	"github.com/becent/golang-common/gin-handler"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()

	cfg := &gin_handler.Config{
		AppName: "MyApp",

		CheckSignature: false,
		// AppId:          "1000",
		// SecretKey:      "XXXXXXXXXXXXXXX",
	}
	engine.Use(gin_handler.NewHandler(cfg))

	engine.Any("/hello", Hello)

	engine.Run(":8080")
}

func Hello(c *gin.Context) {
	h := gin_handler.DefaultHandler(c)

	name := h.StringParam("name")

	h.SuccessResponse("hello " + name)
}
