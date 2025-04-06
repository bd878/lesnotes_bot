package logger

import "go.uber.org/zap"

type Logger struct {
	*zap.SugaredLogger
}

var Log *Logger = (*Logger)(nil)

func init() {
	Log = NewProduction()
}

func NewProduction(options ...zap.Option) *Logger {
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	return &Logger{zap.Must(config.Build(options...)).Sugar()}
}

func NewDevelopment(options ...zap.Option) *Logger {
	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	return &Logger{zap.Must(config.Build(options...)).Sugar()}
}

func NewLog() *Logger {
	return Log
}