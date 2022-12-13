package cfg

func Get[T any](key string) (t T, ok bool) {
	value := defaultStore.Get(key)
	t, ok = value.(T)
	return
}

func Store(key string, value any) {
	defaultStore.Store(key, value)
}
