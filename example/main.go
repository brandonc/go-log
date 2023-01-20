package main

import (
	"errors"

	"github.com/brandonc/go-log"
)

func main() {
	logger := log.NewLoggerFromEnv()

	logger.Debug("This is a debug message")
	logger.Info("This is an info message")
	logger.Warn("This is a warning")
	logger.Errorf("This is an error: %v", errors.New("some error"))
}
