package main

import (
	"fmt"

	"github.com/yarikTri/darty"
	"github.com/yarikTri/darty/pkg/handler"
	"github.com/yarikTri/darty/pkg/repository"
	"github.com/yarikTri/darty/pkg/service"

	"github.com/spf13/viper" // config reader
	"go.uber.org/zap"        // logger
)

const port = "8080"

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("Zap logger can't be defined: %+v", err)
	}

	if err = initConfig(); err != nil {
		logger.Error("Error occured while initing configs" + err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(darty.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logger.Error("Error occured while launching server: " + err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
