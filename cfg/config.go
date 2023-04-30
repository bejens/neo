package cfg

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/bejens/neo/cfg/envx"
	"github.com/bejens/neo/cfg/filex"
	"github.com/bejens/neo/cfg/parser"
)

var config = &Config{
	Storage: defaultStore,
	parsers: []parser.Parser{
		&filex.YamlParser{
			Path: "app.yaml",
		},
		&envx.EnvParser{
			Prefix: "neo",
			Seg:    ".",
			Sep:    "=",
		},
	},
}

func InitCfg() error {

	for _, p := range config.parsers {
		m, err := p.Parse()
		if err != nil {
			return err
		}
		if err := config.Storage.Merge(m); err != nil {
			return err
		}
	}

	profile, ok := Get[string]("profile")
	if !ok {
		return nil
	}

	path := fmt.Sprintf("app_%s.yaml", profile)
	p := filex.YamlParser{Path: path}
	m, err := p.Parse()
	if err != nil {
		return err
	}
	if err := config.Storage.Merge(m); err != nil {
		return err
	}
	return nil
}

type Config struct {
	Storage Storage
	parsers []parser.Parser
}

func Get[T any](key string) (t T, ok bool) {
	value, ok := defaultStore.Get(key)
	if !ok {
		return t, false
	}

	v, ok := value.(T)
	if ok {
		return v, ok
	}
	vs := fmt.Sprintf("%s", value)
	switch interface{}(t).(type) {
	case int:
		v1, err := strconv.ParseInt(vs, 10, 64)
		if err != nil {
			return t, false
		}
		return interface{}(int(v1)).(T), true
	case int8:
		v1, err := strconv.ParseInt(vs, 10, 64)
		if err != nil {
			return t, false
		}
		return interface{}(int8(v1)).(T), true
	case int16:
		v1, err := strconv.ParseInt(vs, 10, 64)
		if err != nil {
			return t, false
		}
		return interface{}(int16(v1)).(T), true
	case int32:
		v1, err := strconv.ParseInt(vs, 10, 64)
		if err != nil {
			return t, false
		}
		return interface{}(int32(v1)).(T), true
	case int64:
		v1, err := strconv.ParseInt(vs, 10, 64)
		if err != nil {
			return t, false
		}
		return interface{}(v1).(T), true
	case uint:
		v1, err := strconv.ParseUint(vs, 10, 64)
		if err != nil {
			return t, false
		}
		return interface{}(uint(v1)).(T), true
	case uint8:
		v1, err := strconv.ParseUint(vs, 10, 64)
		if err != nil {
			return t, false
		}
		return interface{}(uint8(v1)).(T), true
	case uint16:
		v1, err := strconv.ParseUint(vs, 10, 64)
		if err != nil {
			return t, false
		}
		return interface{}(uint16(v1)).(T), true
	case uint32:
		v1, err := strconv.ParseUint(vs, 10, 64)
		if err != nil {
			return t, false
		}
		return interface{}(uint32(v1)).(T), true
	case uint64:
		v1, err := strconv.ParseUint(vs, 10, 64)
		if err != nil {
			return t, false
		}
		return interface{}(v1).(T), true
	case float64:
		v1, err := strconv.ParseFloat(vs, 64)
		if err != nil {
			return t, false
		}
		return interface{}(v1).(T), true
	case float32:
		v1, err := strconv.ParseFloat(vs, 64)
		if err != nil {
			return t, false
		}
		return interface{}(float32(v1)).(T), true
	case bool:
		v1, err := strconv.ParseBool(vs)
		if err != nil {
			return t, false
		}
		return interface{}(v1).(T), true
	case string:
		return interface{}(vs).(T), true
	default:
		bs, err := json.Marshal(value)
		if err != nil {
			return t, false
		}
		if err := json.Unmarshal(bs, &t); err == nil {
			return t, true
		}
		return t, false
	}
}

func Store(key string, value any) {
	defaultStore.Store(key, value)
}
