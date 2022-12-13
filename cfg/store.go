package cfg

import "sync"

var defaultStore *localStore

type localStore struct {
	m sync.Map
}

func (store *localStore) Store(key string, value any) {
	store.m.Store(key, value)
}

func (store *localStore) Get(key string) any {
	value, ok := store.m.Load(key)
	if ok {
		return value
	}
	return nil
}

type Storage interface {
	Store(key string, value any)
	Get(key string) any
}
