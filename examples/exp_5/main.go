package main

import (
	"log/slog"
	"os"
)

func main() {
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	log := slog.New(h).With("App", "exp")

	log.Debug("Debug message")
	log.Info("Info message", "request_id", 1)
	log.Warn("Warn message")
	log.Error("Error message", "request_id", 1)
}
