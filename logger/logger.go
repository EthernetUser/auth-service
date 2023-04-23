package logger

import (
	"time"

	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func InitLogger() {
	logger.SetFormatter(&logrus.TextFormatter{ForceColors: true, TimestampFormat: time.DateTime, FullTimestamp: true})
	logger.SetOutput(colorable.NewColorableStdout())
}

func Info(args ...any) {
	logger.Info(args...)
}

func Infof(format string, v ...any) {
	logger.WithTime(time.Now()).Infof(format, v...)
}

func Fatal(v ...any) {
	logger.Fatal(v...)
}

func Error(v ...any) {
	logger.Error(v...)
}
