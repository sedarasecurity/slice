package slice

import (
	"sort"
	"strings"
)

type StringSlice map[string]uint8

// type StringSlice struct {
// 	lock sync.Mutex
// 	s    map[string]struct{}
// }

func NewStringSlice() StringSlice {
	return make(map[string]uint8)
}

// Add adds the element to the slice
func (slice StringSlice) Add(keys ...string) {
	for _, key := range keys {
		key = strings.TrimSpace(key)
		if len(key) > 0 {
			if slice.Has(key) {
				continue
			}
			slice[key] = 0
		}
	}
}

// Remove removes the element from the slice
func (slice StringSlice) Remove(key string) {
	delete(slice, key)
}

// Has checks for the element in the slice
func (slice StringSlice) Has(key string) bool {
	for k := range slice {
		if strings.EqualFold(key, k) {
			return true
		}
	}
	return false
}

// Contains checks for the substring element in all of the values in the slice
func (slice StringSlice) Contains(substr string) bool {
	for _, value := range slice.Sorted() {
		haystack := strings.ToLower(value)
		needle := strings.ToLower(substr)
		if strings.Contains(haystack, needle) {
			return true
		}
	}
	return false
}

// Copy creates a copy of the string slice and returns it
func (slice StringSlice) Copy() StringSlice {
	cp := NewStringSlice()

	if slice == nil {
		return cp
	}

	cp.Add(slice.Sorted()...)

	return cp
}

// Reset empties the container
func (slice StringSlice) Reset() {
	for _, v := range slice.keys() {
		delete(slice, v)
	}
}

// Sorted returns a sorted string slice
func (slice StringSlice) Sorted() []string {
	arr := slice.keys()
	sort.Strings(arr)

	return arr
}

func (slice StringSlice) keys() []string {
	var arr []string

	if slice == nil {
		return arr
	}

	for key := range slice {
		arr = append(arr, key)
	}
	return arr
}
