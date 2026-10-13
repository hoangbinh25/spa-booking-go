// @title Spa Booking API
// @version 1.0
// @description API for the Spa Booking System.
// @host localhost:8080
// @BasePath /api/v1
// @schemes http

package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/hoangbinh25/spa-booking/internal/branch"
	"github.com/hoangbinh25/spa-booking/internal/config"
	"github.com/hoangbinh25/spa-booking/internal/database"
	"github.com/hoangbinh25/spa-booking/internal/health"
	"github.com/hoangbinh25/spa-booking/internal/httpapi"

	_ "github.com/hoangbinh25/spa-booking/docs"
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

	healthHandler := health.NewHandler(db)
	branchRepository := branch.NewRepository(db)
	branchHandler := branch.NewHandler(branchRepository, branchRepository)

	router := httpapi.NewRouter(httpapi.Handlers{
		Health: healthHandler,
		Branch: branchHandler,
	})

	address := fmt.Sprintf(":%d", cfg.AppPort)
	if err := router.Run(address); err != nil {
		log.Fatalf("Run HTTP server: %v", err)
	}
}
