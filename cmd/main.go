package main

import (
	"fmt"
	"os" // operating system

	"github.com/joho/godotenv"
	"github.com/yarikTri/darty"
	"github.com/yarikTri/darty/pkg/handler"
	"github.com/yarikTri/darty/pkg/repository"
	"github.com/yarikTri/darty/pkg/service"

	_ "github.com/lib/pq"    // postgres driver
	"github.com/spf13/viper" // config reader
	"go.uber.org/zap"        // logger
)

func main() {
	// zap logger
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("Zap logger can't be defined: %+v", err)
	}

	// viper and godotenv configs
	if err = initConfig(); err != nil {
		logger.Error("Error while initing configs: " + err.Error())
	}
	if err = godotenv.Load(); err != nil {
		logger.Error("Error while loading env vars: " + err.Error())
	}

	// connect to db
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		logger.Error("Error while connecting to repository: " + err.Error())
	}

	// architecture
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// server
	srv := new(darty.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logger.Error("Error while launching server: " + err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
