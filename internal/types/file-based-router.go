package types

import (
	"html/template"
	"os"
	"strings"

	"github.com/AYGA2K/go-file-based-router/internal/constants"
	"github.com/AYGA2K/go-file-based-router/internal/util"
)

type FileBasedRouter struct {
	routesCode   string
	handlersCode string
}

func NewFileBasedRouter() FileBasedRouter {
	return FileBasedRouter{
		routesCode:   constants.NOFILESROUTETEMPLATE,
		handlersCode: constants.NOFILESHANDLERTEMPLATE,
	}
}

func (f *FileBasedRouter) scan(path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return err
	}
	for _, v := range files {
		fileName := v.Name()
		_, err := template.ParseFiles(constants.PAGESPATH + fileName)
		if err != nil {
			return err
		}
		name := strings.Split(fileName, ".")[0]
		handler := name + "Handler"
		routeCode := util.GenerateRegisterRoute(name, handler)
		f.routesCode = f.routesCode + routeCode + "\n"
		handlerCode := util.GenerateHandler(name)
		f.handlersCode = f.handlersCode + handlerCode + "\n"
	}
	return nil
}

func (f *FileBasedRouter) CreateRoutesFile(path string) error {
	err := f.scan(path)
	if err != nil {
		return err
	}
	err = util.CreateFile(constants.ROUTESFILEPATH)
	if err != nil {
		return err
	}

	err = util.WriteToRoutesFile(constants.ROUTESFILEPATH, f.routesCode, f.handlersCode)
	if err != nil {
		return err
	}

	return nil
}
