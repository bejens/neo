package zapx

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap/zapcore"
)

func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		LevelKey:            "level",
		TimeKey:             "time",
		NameKey:             "logger",
		CallerKey:           "caller",
		MessageKey:          "msg",
		StacktraceKey:       "stack",
		LineEnding:          zapcore.DefaultLineEnding,
		EncodeLevel:         zapcore.LowercaseLevelEncoder,
		EncodeTime:          zapcore.ISO8601TimeEncoder,
		EncodeDuration:      zapcore.SecondsDurationEncoder,
		EncodeCaller:        zapcore.ShortCallerEncoder,
		NewReflectedEncoder: newReflectedEncoder,
	}
}

// newReflectedEncoder replace encoding/json to jsoniter
func newReflectedEncoder(w io.Writer) zapcore.ReflectedEncoder {
	enc := jsoniter.NewEncoder(w)
	// For consistency with our custom JSON encoder.
	enc.SetEscapeHTML(false)
	return enc
}
