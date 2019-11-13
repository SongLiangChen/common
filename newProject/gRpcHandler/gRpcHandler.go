package gRpcHandler

import (
	"os"
)

func G_gRpcHandler(projectName string) error {
	if err := os.Mkdir(projectName+"/gRpcHandler", 755); err != nil {
		return err
	}

	file, err := os.OpenFile(projectName+"/gRpcHandler/example.go", os.O_CREATE|os.O_RDWR, 755)
	if err != nil {
		return err
	}

	if _, err := file.WriteString(gRpcHandler_temple); err != nil {
		return err
	}

	return nil
}

var gRpcHandler_temple = `package gRpcHandler

import (
	"fmt"

	"github.com/becent/golang-common/grpc-end"
)

// a example for grpc request
func Hello(c *grpc_end.GRpcContext) {
	// init params here
	name := c.StringParamDefault("name", "Tom")
	age := c.IntParam("age")

	// do service here
	// ...

	c.SuccessResponse(fmt.Sprintf("hello %v years old %v", age, name))
}
`
