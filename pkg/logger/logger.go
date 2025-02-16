package logger

import (
	"log/slog"
	"os"
)

func SetupLogger(env string) *slog.Logger {
	var logLevel slog.Level
	if env == "development" {
		logLevel = slog.LevelDebug
	} else {
		logLevel = slog.LevelInfo
	}

	handlerOptions := &slog.HandlerOptions{
		Level: logLevel,
	}

	handler := slog.NewTextHandler(os.Stdout, handlerOptions)
	logger := slog.New(handler)

	return logger
}
