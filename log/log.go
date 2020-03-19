package log

import (
	"github.com/sirupsen/logrus"
)

type Logger struct {
	Log *logrus.Logger
}

type Fields = logrus.Fields

func NewLogger() *Logger {
	logger := &Logger{}

	logger.Log = logrus.New()

	return logger
}

// func (l *Logger) Debug(fields ...Field) {
// 	l.log.De(fields...)
// }

// func (l *Logger) Info(message string, fields ...Field) {
// 	l.Info(message, fields...)
// }

// func (l *Logger) Warn(message string, fields ...Field) {
// 	l.Warn(message, fields...)
// }

// func (l *Logger) Error(message string, fields ...Field) {
// 	l.Error(message, fields...)
// }
