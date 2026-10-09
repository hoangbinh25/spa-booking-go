package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/hoangbinh25/spa-booking/internal/config"
	"github.com/hoangbinh25/spa-booking/internal/database"
	"github.com/hoangbinh25/spa-booking/internal/health"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("invalid application configuration: %v", err)
	}

	slog.Info(
		"starting application",
		"environment", cfg.AppEnv,
		"port", cfg.AppPort,
	)

	ctx := context.Background()

	db, err := database.NewPostgresPool(ctx, cfg.DatabaseURL())
	if err != nil {
		log.Fatalf("connect to database: %v", err)
	}

	defer db.Close()

	slog.Info("database connected")

	router := gin.New()
	router.Use(
		gin.Logger(),
		gin.Recovery(),
	)

	healthHandler := health.NewHandler(db)
	router.GET("/health", healthHandler.Check)

	address := fmt.Sprintf(":%d", cfg.AppPort)
	if err := router.Run(address); err != nil {
		log.Fatalf("Run HTTP server: %v", err)
	}
}
