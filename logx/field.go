package logx

import "github.com/bejens/neo/logx/logger"

func Int(key string, value int) logger.Field {
	return logger.Field{
		Key:   key,
		Value: value,
		Type:  logger.Int,
	}
}

func Int8(key string, value int8) logger.Field {
	return logger.Field{
		Key:   key,
		Value: value,
		Type:  logger.Int8,
	}
}

func Int16(key string, value int16) logger.Field {
	return logger.Field{
		Key:   key,
		Value: value,
		Type:  logger.Int16,
	}
}

func Int32(key string, value int32) logger.Field {
	return logger.Field{
		Key:   key,
		Value: value,
		Type:  logger.Int32,
	}
}

func Int64(key string, value int64) logger.Field {
	return logger.Field{
		Key:   key,
		Value: value,
		Type:  logger.Int64,
	}
}

func Uint(key string, value uint) logger.Field {
	return logger.Field{
		Key:   key,
		Value: value,
		Type:  logger.Uint,
	}
}

func Uint8(key string, value uint8) logger.Field {
	return logger.Field{
		Key:   key,
		Value: value,
		Type:  logger.Uint8,
	}
}

func Uint16(key string, value uint16) logger.Field {
	return logger.Field{
		Key:   key,
		Value: value,
		Type:  logger.Uint16,
	}
}

func Uint32(key string, value uint32) logger.Field {
	return logger.Field{
		Key:   key,
		Value: value,
		Type:  logger.Uint32,
	}
}

func Uint64(key string, value uint64) logger.Field {
	return logger.Field{
		Key:   key,
		Value: value,
		Type:  logger.Uint64,
	}
}

func Float32(key string, value float32) logger.Field {
	return logger.Field{
		Key:   key,
		Value: value,
		Type:  logger.Float32,
	}
}

func Float64(key string, value float64) logger.Field {
	return logger.Field{
		Key:   key,
		Value: value,
		Type:  logger.Float64,
	}
}

func String(key, value string) logger.Field {
	return logger.Field{
		Key:   key,
		Value: value,
		Type:  logger.String,
	}
}

func Any[T any](key string, value T) logger.Field {
	return logger.Field{
		Key:   key,
		Value: value,
		Type:  logger.Any,
	}
}

func Slice[T comparable](key string, value []T) logger.Field {
	return logger.Field{
		Key:   key,
		Value: value,
		Type:  logger.Slice,
	}
}

func Map[K comparable, V any](key string, value map[K]V) logger.Field {
	return logger.Field{
		Key:   key,
		Value: value,
		Type:  logger.Map,
	}
}

func Err(err error) logger.Field {
	return logger.Field{
		Key:   "error",
		Value: err.Error(),
		Type:  logger.Error,
	}
}
