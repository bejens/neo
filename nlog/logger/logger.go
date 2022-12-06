package logger

type Logger interface {
	Log(level Level, v ...any)
}
