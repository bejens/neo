package cfg

import (
	"strings"
)

var defaultStore *localStore

type localStore struct {
	m map[string]any
}

func (store *localStore) Store(key string, value any) {
	store.m[key] = value
}

func (store *localStore) Get(key string) (any, bool) {
	keys := strings.Split(key, ".")
	return store.get(store.m, keys...)
}

func (store *localStore) get(m map[string]any, keys ...string) (any, bool) {

	if len(keys) == 0 {
		return "", false
	}

	v, ok := m[keys[0]]
	if !ok {
		return "", false
	}

	if len(keys) == 1 {
		return v, true
	}

	m1, ok := v.(map[string]any)
	if !ok {
		return "", false
	}

	return store.get(m1, keys[1:]...)
}

type Storage interface {
	Store(key string, value any)
	Get(key string) any
}
