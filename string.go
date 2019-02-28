package slice

import (
	"sort"
	"strings"
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
		key = strings.TrimSpace(key)
		if len(key) > 0 {
			s[key] = struct{}{}
		}
	}
}

// Remove removes the element from the slice
func (s StringSlice) Remove(key string) {
	delete(s, key)
}

// Contains checks for the element in the slice
func (s StringSlice) Contains(key string) bool {
	_, ok := s[key]
	return ok
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
	var (
		keys []string
	)
	for k := range s {
		keys = append(keys, k)
	}
	for _, v := range keys {
		delete(s, v)
	}
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
