package providers

import (
	"go.uber.org/zap"
)

// NewLogger creates a new structured logger
func NewLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic("Failed to create logger: " + err.Error())
	}
	return logger
}
