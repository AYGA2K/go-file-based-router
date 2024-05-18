package util

import (
	"fmt"

	"github.com/AYGA2K/go-file-based-router/internal/constants"
)

func GenerateRegisterRoute(path string, handler string) string {
	registerRouteCode := fmt.Sprintf(constants.ROUTETEMPLATE, path, handler)
	return registerRouteCode
}

func GenerateHandler(name string) string {
	handlerCode := fmt.Sprintf(constants.HANDLERTEMPLATE, name, name)
	return handlerCode
}
