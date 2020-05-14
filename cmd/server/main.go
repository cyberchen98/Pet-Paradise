package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"pet-paradise/api"
	"pet-paradise/config"
	"pet-paradise/log"
	"time"
)

func main() {
	fmt.Println(VERSION)
	fmt.Println("Pet-Paradise Server")
	if err := config.ParseConfig("config.yaml"); err != nil {
		os.Exit(1)
	}
	log.Logger().Info("start server")

	srv := &http.Server{
		Addr:    config.Server.API,
		Handler: api.InitRouter(),
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Logger().Fatal("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Logger().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Logger().Fatal("Server Shutdown:", err)
	}
	log.Logger().Info("Server exiting")
}
