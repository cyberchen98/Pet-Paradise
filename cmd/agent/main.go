package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pet-paradise/middleware"
	"time"
)

func main() {
	if err := ParseConfig("config.yaml"); err != nil {
		os.Exit(1)
	}
	r := gin.Default()
	base := r.Group("/file")

	base.GET("/", Get)
	base.DELETE("/", Delete)

	admin := r.Group("/admin")
	authFunc := middleware.AuthMiddleware()
	adminAuthFunc := middleware.AdminAuthMiddleware()
	admin.Use(authFunc, adminAuthFunc)

	admin.POST("/upload", Upload)
	admin.GET("/list", List)

	srv := &http.Server{
		Addr:    Agent.Port,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Agent ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Agent Shutdown:", err)
	}
}
