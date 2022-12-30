package logx

import (
	"github.com/bejens/neo/logx/logger"
	"github.com/bejens/neo/logx/zapx"
)

var loggerCore = zapx.Default()

func WarpLogger(logger logger.Logger) {
	loggerCore = logger
}

func Debug(msg string, args ...any) {
	loggerCore.Log(logger.DebugLevel, msg, args...)
}

func Warn(msg string, args ...any) {
	loggerCore.Log(logger.WarnLevel, msg, args...)
}

func Error(msg string, args ...any) {
	loggerCore.Log(logger.ErrorLevel, msg, args...)
}

func Info(msg string, args ...any) {
	loggerCore.Log(logger.InfoLevel, msg, args...)
}

func Fatalln(msg string, args ...any) {
	loggerCore.Log(logger.FatalLevel, msg, args...)
}

func Panic(msg string, args ...any) {
	loggerCore.Log(logger.PanicLevel, msg, args...)
}
