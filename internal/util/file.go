package util

import (
	"fmt"
	"os"

	"github.com/AYGA2K/go-file-based-router/internal/constants"
)

func CreateFile(path string) error {
	_, err := os.Create(path)
	if err != nil {
		return err
	}
	return nil
}

func WriteToRoutesFile(path string, routes string, handlers string) error {
	file, err := os.OpenFile(path, os.O_WRONLY, 0444)
	if err != nil {
		return err
	}
	defer file.Close()
	content := fmt.Sprintf(constants.ROUTESFILETEMPLATE, routes, handlers)
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}
