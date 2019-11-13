package handler

import (
	"os"
)

func G_handler(projectName string) error {
	if err := os.Mkdir(projectName+"/handler", 755); err != nil {
		return err
	}

	file, err := os.OpenFile(projectName+"/handler/example.go", os.O_CREATE|os.O_RDWR, 755)
	if err != nil {
		return err
	}

	if _, err := file.WriteString(handler_temple); err != nil {
		return err
	}

	return nil
}

var handler_temple = `package handler

import (
	"fmt"

	"github.com/becent/golang-common/gin-handler"
	"github.com/gin-gonic/gin"
)

// a example for gin request
func Hello(c *gin.Context) {
	h := gin_handler.DefaultHandler(c)

	// init the request' param here
	var (
		name = h.StringParam("name")
		age  = h.IntParam("age")
	)

	// do service here
	// ...

	h.SuccessResponse(fmt.Sprintf("hello %v years old %v", age, name))
}
`
