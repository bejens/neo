package zapx

import (
	"fmt"

	"github.com/bejens/neo/logx/logger"

	"go.uber.org/zap"
)

func Default() logger.Logger {
	zl, err := zap.NewProduction(zap.AddCallerSkip(2))
	if err != nil {
		panic(err)
	}

	return &ZLogger{logger: zl}
}

type ZLogger struct {
	logger *zap.Logger
}

func (zl *ZLogger) Log(level logger.Level, msg string, args ...any) {

	fields := zl.assert(args...)

	switch level {
	case logger.InfoLevel:
		zl.logger.Info(msg, fields...)
	case logger.DebugLevel:
		zl.logger.Debug(msg, fields...)
	case logger.WarnLevel:
		zl.logger.Warn(msg, fields...)
	case logger.ErrorLevel:
		zl.logger.Error(msg, fields...)
	case logger.FatalLevel:
		zl.logger.Fatal(msg, fields...)
	case logger.PanicLevel:
		zl.logger.Panic(msg, fields...)
	default:
		zl.logger.Warn(msg, fields...)
	}
}

func (zl *ZLogger) assert(args ...any) (fields []zap.Field) {
	for index, arg := range args {
		if field, ok := arg.(zap.Field); ok {
			fields = append(fields, field)
		} else {
			fields = append(fields, zap.Any(fmt.Sprintf("arg%d", index), arg))
		}
	}
	return
}

func NewLogger() (logger.Logger, error) {
	zl, err := zap.NewProduction(zap.AddCallerSkip(2))
	if err != nil {
		return nil, err
	}

	return &ZLogger{logger: zl}, nil
}
