package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	httpAdapter "github.com/bakhtybayevn/powerbook/internal/adapters/http"
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

// @host localhost:8080
// @BasePath /
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Printf("Starting powerbook in %s mode on port %d\n", cfg.App.Environment, cfg.App.Port)

	router := gin.Default()

	// initialize HTTP server (handlers)
	srv := httpAdapter.NewServer(router, cfg)
	srv.RegisterRoutes()

	// start server
	addr := fmt.Sprintf(":%d", cfg.App.Port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
