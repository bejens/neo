package zapx

import (
	"errors"

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

	fields, err := zl.assert(args...)
	if err != nil {
		zl.logger.Error(err.Error())
	}

	switch level {
	case logger.InfoLevel:
		zl.logger.Info(msg, fields...)
	case logger.DebugLevel:
		zl.logger.Debug(msg, fields...)
	case logger.WarnLevel:
		zl.logger.Warn(msg, fields...)
	case logger.ErrorLevel:
		zl.logger.Error(msg, fields...)
	case logger.PanicLevel:
		zl.logger.Panic(msg, fields...)
	default:
		zl.logger.Warn(msg, fields...)
	}
}

func (zl *ZLogger) assert(args ...any) (fields []zap.Field, err error) {
	for _, arg := range args {
		if field, ok := arg.(zap.Field); ok {
			fields = append(fields, field)
		} else {
			err = errors.New("only zap.Field can be use as args")
			return
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
