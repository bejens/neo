package envx

import (
	"errors"
	"os"
	"strings"
)

type EnvParser struct {
	prefix string
	seg    string
	sep    string
}

func (ep *EnvParser) Parse() (m map[string]any, err error) {

	envs := make(map[string]any)

	environ := os.Environ()
	for _, env := range environ {
		if !strings.HasPrefix(env, ep.prefix) {
			continue
		}
		env = strings.TrimPrefix(env, ep.prefix)
		slice := strings.SplitN(env, ep.sep, 2)
		if len(slice) < 2 {
			continue
		}
		if err := ep.store(envs, strings.Split(slice[0], ep.seg), slice[1]); err != nil {
			return m, err
		}
	}

	return envs, nil
}

func (ep *EnvParser) store(m map[string]any, keys []string, value string) error {
	if len(keys) == 1 {
		m[keys[0]] = value
		return nil
	}

	key := keys[0]
	m1, ok := m[key]
	if !ok {
		m1 = map[string]any{}
		m[key] = m1
	}

	m2, ok := m1.(map[string]any)
	if !ok {
		return errors.New("type Conversion Error")
	}

	return ep.store(m2, keys[1:], value)
}
