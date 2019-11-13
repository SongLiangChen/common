package service

import (
	"os"
)

func G_service(projectName string) error {
	if err := os.Mkdir(projectName+"/service", 755); err != nil {
		return err
	}

	return nil
}
