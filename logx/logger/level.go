package logger

type Level int8

const (
	DebugLevel Level = iota - 1
	WarnLevel
	ErrorLevel
	InfoLevel
	PanicLevel
)
