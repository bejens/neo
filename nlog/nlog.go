package nlog

import (
	"github.com/bejens/neo/nlog/logger"
	"github.com/bejens/neo/nlog/zap"
)

var defaultLogger logger.Logger = zap.Default()

func Info(msg string) {
	defaultLogger.Log(msg)
}
