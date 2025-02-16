package main

import (
	"log/slog"
	"os"

	"merch-shop/internal/app"
	"merch-shop/internal/config"
	"merch-shop/pkg/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.SetupLogger(cfg.App.Env)

	log.Info("Starting application")

	application, err := app.New(log)
	if err != nil {
		log.Error("Failed to initialize application", slog.String("error", err.Error()))
		os.Exit(1)
	}

	if err := application.HTTPServer.Run(":8080"); err != nil {
		log.Error("Error running server", slog.String("error", err.Error()))
		os.Exit(1)
	}

	log.Info("Server started successfully")
}
