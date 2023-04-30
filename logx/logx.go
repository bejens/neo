package logx

import (
	"github.com/bejens/neo/logx/logger"
	"github.com/bejens/neo/logx/zapx"
)

var core = zapx.Default()

func Warp(logger logger.Logger) {
	core = logger
}

func Debug(msg string, args ...logger.Field) {
	core.Log(logger.DebugLevel, msg, args...)
}

func Warn(msg string, args ...logger.Field) {
	core.Log(logger.WarnLevel, msg, args...)
}

func Error(msg string, args ...logger.Field) {
	core.Log(logger.ErrorLevel, msg, args...)
}

func Info(msg string, args ...logger.Field) {
	core.Log(logger.InfoLevel, msg, args...)
}

func Fatal(msg string, args ...logger.Field) {
	core.Log(logger.FatalLevel, msg, args...)
}

func Panic(msg string, args ...logger.Field) {
	core.Log(logger.PanicLevel, msg, args...)
}

func Sync() error {
	return core.Sync()
}
