package cfg

func InitCfg() {

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
