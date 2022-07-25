package main

import (
	"log"

	"github.com/yarikTri/darty"
	"go.uber.org/zap"
)

const port = "8080"

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Zap logger can't be defined: %+v", err)
	}

	srv := new(darty.Server)
	if err := srv.Run(port); err != nil {
		logger.Error("Error occured while launching server")
	}
}
