package cache

type CacheData struct {
	Value any
	Expires  int32
}

var cache map[string]CacheData

func Init() {
	cache = make(map[string]CacheData, 0)
}

func Set(key string, value any, expires int32) {
	cache[key] = CacheData{
		Value: value,
		Expires: expires,
	}
}

func Get(key string) (any, bool) {
	data, ok := cache[key]
	return data.Value, ok
}

func Delete(key string) {
	delete(cache, key)
}


