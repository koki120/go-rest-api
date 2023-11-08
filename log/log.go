package log

import (
	"log/slog"
	"os"

	"github.com/koki120/go-rest-api/config"
)

var (
	logger *slog.Logger
)

func NewLogger() *slog.Logger {
	if logger != nil {
		return logger
	}
	if config.IsDevelopment() {
		logger = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}
	return logger
}
