package envx

import (
	"errors"
	"io"
	"os"
	"strings"
)

type EnvParser struct {
	prefix string
	seg    string
	sep    string
}

func (ep *EnvParser) Parse(_reader io.Reader) (m map[string]any, err error) {

	envs := make(map[string]any)

	environ := os.Environ()
	for _, env := range environ {
		if !strings.HasPrefix(env, ep.prefix) {
			continue
		}
		envSlice := strings.Split(env, ep.sep)
		if len(envSlice) < 2 {
			continue
		}
		keys := strings.Split(envSlice[0], ep.seg)
		if err := ep.store(envs, keys, envSlice[1]); err != nil {
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

	m1, ok := m[keys[0]]
	if !ok {
		m1 = map[string]any{}
		m[keys[0]] = m1
	}

	m2, ok := m1.(map[string]any)
	if !ok {
		return errors.New("type Conversion Error")
	}

	return ep.store(m2, keys[1:], value)
}
