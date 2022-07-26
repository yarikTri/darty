package main

import (
	"fmt"

	"github.com/yarikTri/darty"
	"github.com/yarikTri/darty/pkg/handler"
	"github.com/yarikTri/darty/pkg/repository"
	"github.com/yarikTri/darty/pkg/service"

	"go.uber.org/zap" // logger
)

const port = "8080"

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("Zap logger can't be defined: %+v", err)
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(darty.Server)
	if err := srv.Run(port, handlers.InitRoutes()); err != nil {
		logger.Error("Error occured while launching server: " + err.Error())
	}
}
