package logger

import "go.uber.org/zap"

func NewLogger() (*zap.Logger, error) {
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout", "app.log"}

	return config.Build()
}