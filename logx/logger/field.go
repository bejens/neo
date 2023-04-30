package logger

type Field struct {
	Key   string
	Value any
	Type  Type
}

type Type int

const (
	Int Type = iota
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Float32
	Float64
	String
	Any
	Map
	Slice
	Error
)
