package main

import (
	"fmt"
	"log"

	"github.com/AYGA2K/go-file-based-router/internal/constants"
	"github.com/AYGA2K/go-file-based-router/internal/server"
	"github.com/AYGA2K/go-file-based-router/internal/types"
)

func main() {
	router := types.NewFileBasedRouter()
	err := router.CreateRoutesFile(constants.PAGESPATH)
	if err != nil {
		log.Fatal(err.Error())
	}

	server := server.NewServer()

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
