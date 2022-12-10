package logger

type Logger interface {
	Log(level Level, msg string, args ...any)
}
