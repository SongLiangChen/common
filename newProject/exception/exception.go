package exception

import (
	"os"
)

func G_exception(projectName string) error {
	if err := os.Mkdir(projectName+"/exception", 755); err != nil {
		return err
	}

	file, err := os.OpenFile(projectName+"/exception/exception.go", os.O_CREATE|os.O_RDWR, 755)
	if err != nil {
		return err
	}

	if _, err := file.WriteString(exception_temple); err != nil {
		return err
	}

	return nil
}

var exception_temple = `package exception

import (
	"github.com/becent/golang-common/exception"
)

var (
	SERVICE_BUSY = &exception.Exception{
		Code: 1000,
		Message: "服务繁忙",
	}

	// TODO add your exception here...
)
`
