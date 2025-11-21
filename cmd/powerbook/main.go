package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	httpAdapter "github.com/bakhtybayevn/powerbook/internal/adapters/http"
	"github.com/gin-gonic/gin"

	_ "github.com/bakhtybayevn/powerbook/internal/adapters/http/docs"
	"github.com/bakhtybayevn/powerbook/internal/config"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter: Bearer <token>

// @title PowerBook API
// @version 1.0
// @description API for PowerBook reading competition platform

// @BasePath /api/v1
// @schemes https http
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	router := gin.Default()

	server := httpAdapter.NewServer(router, cfg)
	server.RegisterRoutes()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.App.Port),
		Handler: router,
	}

	// Run server in background goroutine
	go func() {
		fmt.Printf("Starting PowerBook API on port %d...\n", cfg.App.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server startup error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	fmt.Println("Shutting down PowerBook API...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Forced shutdown: %v", err)
	}

	fmt.Println("PowerBook API stopped gracefully.")
}
