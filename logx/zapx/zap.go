package zapx

import (
	"os"

	"github.com/bejens/neo/logx/logger"
	"go.uber.org/zap/zapcore"

	"go.uber.org/zap"
)

func Default() logger.Logger {

	zapc := zapcore.NewCore(
		// 编码器配置
		zapcore.NewJSONEncoder(newEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(zap.DebugLevel),
	)

	zl := zap.New(zapc, zap.AddCaller(), zap.AddCallerSkip(2))

	return &ZLogger{logger: zl}
}

func NewLogger() (logger.Logger, error) {
	zapc := zapcore.NewCore(
		// 编码器配置
		zapcore.NewJSONEncoder(newEncoderConfig()),
		zapcore.AddSync(os.Stdout),
		zap.NewAtomicLevelAt(zap.DebugLevel),
	)

	zl := zap.New(zapc, zap.AddCaller(), zap.AddCallerSkip(2))

	return &ZLogger{logger: zl}, nil
}

type ZLogger struct {
	logger *zap.Logger
}

func (zl *ZLogger) Sync() error {
	return zl.logger.Sync()
}

func (zl *ZLogger) Log(level logger.Level, msg string, args ...logger.Field) {

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
		zl.logger.Info(msg, fields...)
	}
}

func (zl *ZLogger) assert(args ...logger.Field) (fields []zap.Field) {
	for _, arg := range args {
		switch arg.Type {
		case logger.Int:
			fields = append(fields, zap.Int(arg.Key, arg.Value.(int)))
		case logger.Int8:
			fields = append(fields, zap.Int8(arg.Key, arg.Value.(int8)))
		case logger.Int16:
			fields = append(fields, zap.Int16(arg.Key, arg.Value.(int16)))
		case logger.Int32:
			fields = append(fields, zap.Int32(arg.Key, arg.Value.(int32)))
		case logger.Int64:
			fields = append(fields, zap.Int64(arg.Key, arg.Value.(int64)))
		case logger.Uint:
			fields = append(fields, zap.Uint(arg.Key, arg.Value.(uint)))
		case logger.Uint8:
			fields = append(fields, zap.Uint8(arg.Key, arg.Value.(uint8)))
		case logger.Uint16:
			fields = append(fields, zap.Uint16(arg.Key, arg.Value.(uint16)))
		case logger.Uint32:
			fields = append(fields, zap.Uint32(arg.Key, arg.Value.(uint32)))
		case logger.Uint64:
			fields = append(fields, zap.Uint64(arg.Key, arg.Value.(uint64)))
		case logger.String:
			fields = append(fields, zap.String(arg.Key, arg.Value.(string)))
		case logger.Slice:
			fields = append(fields, zap.Any(arg.Key, arg.Value))
		case logger.Map:
			fields = append(fields, zap.Any(arg.Key, arg.Value))
		case logger.Any:
			fields = append(fields, zap.Any(arg.Key, arg.Value))
		case logger.Error:
			fields = append(fields, zap.Error(arg.Value.(interface{ Error() string })))
		default:
		}
	}
	return
}
