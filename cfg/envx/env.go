package envx

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type EnvParser struct {
	Prefix string
	Seg    string
	Sep    string
}

func (ep *EnvParser) Ext() string {
	return "env"
}

func (ep *EnvParser) Parse() (map[string]any, error) {

	envs := make(map[string]any)

	environ := os.Environ()
	for _, env := range environ {
		if !strings.HasPrefix(env, ep.Prefix) {
			continue
		}
		env = strings.TrimPrefix(env, ep.Prefix+ep.Seg)
		slice := strings.SplitN(env, ep.Sep, 2)
		if len(slice) < 2 {
			continue
		}
		if err := ep.store(envs, strings.Split(slice[0], ep.Seg), slice[1]); err != nil {
			return envs, err
		}
	}

	return envs, nil
}

func (ep *EnvParser) store(m map[string]any, keys []string, value string) error {
	key := keys[0]
	k, idx, err := ep.isArray(key)
	if err != nil {
		return err
	}
	if idx > -1 {
		return ep.storeSlice(m, k, idx, keys, value)
	}

	if len(keys) == 1 {
		m[key] = value
		return nil
	}

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

func (ep *EnvParser) isArray(key string) (k string, index int, err error) {
	isArray, err := regexp.MatchString(".*(\\[\\d+])$", key)
	if err != nil {
		return key, -1, err
	}
	if !isArray {
		return key, -1, nil
	}

	var (
		begin int
	)
	for i := len(key) - 2; i > 0; i-- {
		if key[i] == '[' {
			begin = i
		}
	}
	idx, err := strconv.ParseInt(key[begin+1:len(key)-1], 10, 64)
	if err != nil {
		return key, -1, err
	}

	return key[:begin], int(idx), nil
}

func (ep *EnvParser) storeSlice(m map[string]any, k string, idx int, keys []string, value string) error {
	v, ok := m[k]
	if !ok {
		if len(keys) == 1 {
			slice := make([]string, idx+1)
			slice[idx] = value
			m[k] = slice
			return nil
		}
		v = []map[string]any{}
		m[k] = v
	}
	if len(keys) == 1 {
		v1, ok := v.([]string)
		if !ok {
			return fmt.Errorf("key is %s, find value in config store: %v,it is not []string", keys[0], v)
		}
		if len(v1)-1 < idx {
			appendSlice := make([]string, idx+1-len(v1))
			appendSlice[len(appendSlice)-1] = value
			m[k] = append(v1, appendSlice...)
			return nil
		}
		v1[idx] = value
		m[k] = v1
		return nil
	}
	v1, ok := v.([]map[string]any)
	if !ok {
		return fmt.Errorf("key is %v, find value in config store: %v,it is not []map[string]any", keys, v)
	}
	if len(v1)-1 < idx {
		appendSlice := make([]map[string]any, idx+1-len(v1))
		item := make(map[string]any)
		err := ep.store(item, keys[1:], value)
		if err != nil {
			return err
		}
		appendSlice[len(appendSlice)-1] = item
		m[k] = append(v1, appendSlice...)
		return nil
	}
	item := v1[idx]
	err := ep.store(item, keys[1:], value)
	if err != nil {
		return err
	}
	v1[idx] = item
	m[k] = v1
	return nil
}
