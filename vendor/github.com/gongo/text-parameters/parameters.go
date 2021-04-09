package parameters

import (
	"sort"
)

type TextParameters map[string]string

func (t TextParameters) Get(key string) string {
	return t[key]
}

func (t TextParameters) Set(key, value string) {
	t[key] = value
}

func (t TextParameters) Del(key string) {
	delete(t, key)
}

func (t TextParameters) Keys() []string {
	var keys []string
	for k := range t {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	return keys
}
