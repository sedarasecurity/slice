package slice

import (
	"sort"
)

type IntSlice map[int]struct{}

// type StringSlice struct {
// 	lock sync.Mutex
// 	s    map[int]struct{}
// }

func NewIntSlice() IntSlice {
	return make(map[int]struct{}, 0)
}

// Add adds the element to the slice
func (s IntSlice) Add(keys ...int) {
	for _, key := range keys {
		s[key] = struct{}{}
	}
}

// Remove removes the element from the slice
func (s IntSlice) Remove(key int) {
	delete(s, key)
}

// Contains checks for the element in the slice
func (s IntSlice) Contains(key int) bool {
	_, ok := s[key]
	return ok
}

func (s IntSlice) Copy() IntSlice {
	cp := NewIntSlice()

	if s == nil {
		return cp
	}

	for k := range s {
		cp.Add(k)
	}

	return cp
}

func (s IntSlice) Reset() {
	var (
		keys []int
	)
	for k := range s {
		keys = append(keys, k)
	}
	for _, v := range keys {
		delete(s, v)
	}
}

// Sorted returns a sorted string slice
func (s IntSlice) Sorted() []int {
	if s == nil {
		return []int{}
	}

	var arr []int

	for key := range s {
		arr = append(arr, key)
	}

	sort.Ints(arr)

	return arr
}
