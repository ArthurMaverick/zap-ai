package logger

import (
	"log/slog"
	"os"
)

// Log is a function to log the messages
func Log() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	slog.SetDefault(logger)
}
