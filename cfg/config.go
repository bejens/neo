package cfg

import (
	"fmt"
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

	t, ok = value.(T)
	return
}

func Store(key string, value any) {
	defaultStore.Store(key, value)
}
