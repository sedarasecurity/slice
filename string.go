package slice

import (
	"sort"
)

type StringSlice map[string]struct{}

// type StringSlice struct {
// 	lock sync.Mutex
// 	s    map[string]struct{}
// }

func NewStringSlice() StringSlice {
	return make(map[string]struct{}, 0)
}

// Add adds the element to the slice
func (s StringSlice) Add(keys ...string) {
	for _, key := range keys {
		s[key] = struct{}{}
	}
}

// Remove removes the element from the slice
func (s StringSlice) Remove(key string) {
	delete(s, key)
}

func (s StringSlice) Copy() StringSlice {
	cp := NewStringSlice()

	if s == nil {
		return cp
	}

	for k := range s {
		cp.Add(k)
	}

	return cp
}

func (s StringSlice) Reset() {
	s = make(map[string]struct{}, 0)
}

// Sorted returns a sorted string slice
func (s StringSlice) Sorted() []string {
	if s == nil {
		return []string{}
	}

	var arr []string

	for key := range s {
		arr = append(arr, key)
	}

	sort.Strings(arr)

	return arr
}
