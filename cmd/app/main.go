package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/danyaobertan/validcard/config"
	"github.com/danyaobertan/validcard/internal/app"
)

// @title Credit Card Validator API
// @version 1.0
// @description This is a simple API to validate credit card information such as card number, expiration month and expiration year

// @host localhost:8080
// @BasePath /
// @Schemes http

func main() {
	if err := config.ReadConfig(); err != nil {
		log.Fatalf("error while reading config: %s", err.Error())
	}

	application := app.NewApp()
	if err := application.Run(); err != nil {
		log.Fatalf("error while running application: %s", err.Error())
	}

	ctx := registerGracefulHandle()
	<-ctx.Done()
}

func registerGracefulHandle() context.Context {
	ctx, cancel := context.WithCancel(context.Background())

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cancel()
	}()

	return ctx
}
