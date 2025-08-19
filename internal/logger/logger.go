package logger

import (
	"log/slog"

	"github.com/natefinch/lumberjack"
)

type Logger struct {
	log *slog.Logger
}

func NewLogger() *Logger {
	lj := &lumberjack.Logger{
		Filename:   "fizzbuzz-api.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     30,
		Compress:   true,
	}
	handler := slog.NewJSONHandler(lj, &slog.HandlerOptions{
		AddSource: true,
	})
	return &Logger{log: slog.New(handler)}
}

func (l *Logger) Info(msg string, args ...any) {
	l.log.Info(msg, args...)
}

func (l *Logger) Error(msg string, args ...any) {
	l.log.Error(msg, args...)
}
