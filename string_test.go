package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSlice_Add(t *testing.T) {
	assert := assert.New(t)
	ss := NewStringSlice()

	ss.Add("a", "b", "c", "a")
	assert.ElementsMatch([]string{"a", "b", "c"}, ss.Sorted())
	ss.Add("a", "b", "c", "a", "", " ", "    ")
	assert.ElementsMatch([]string{"a", "b", "c"}, ss.Sorted())
}

func TestStringSlice_Remove(t *testing.T) {
	assert := assert.New(t)
	ss := NewStringSlice()

	ss.Add("a", "b", "c")
	ss.Remove("c")
	assert.ElementsMatch([]string{"a", "b"}, ss.Sorted())

	ss.Remove("x")
	assert.ElementsMatch([]string{"a", "b"}, ss.Sorted())
}

func TestStringSlice_Has(t *testing.T) {
	assert := assert.New(t)
	ss := NewStringSlice()

	ss.Add("a", "b", "c")
	assert.True(ss.Has("c"))
	assert.False(ss.Has("d"))
}

func TestStringSlice_Contains(t *testing.T) {
	assert := assert.New(t)
	ss := NewStringSlice()

	ss.Add("abcdefg", "hijklmn", "opqrstuv")
	assert.True(ss.Contains("cde"))
	assert.True(ss.Contains("op"))
	assert.True(ss.Contains("tu"))
	assert.False(ss.Contains("wxyz"))
}

func TestStringSlice_Sorted(t *testing.T) {
	assert := assert.New(t)
	ss := NewStringSlice()

	ss.Add("z", "a", "m")
	assert.ElementsMatch([]string{"a", "m", "z"}, ss.Sorted())
}

func TestStringSlice_Reset(t *testing.T) {
	assert := assert.New(t)
	ss := NewStringSlice()

	ss.Add("z", "a", "m")
	assert.ElementsMatch([]string{"a", "m", "z"}, ss.Sorted())

	ss.Reset()
	assert.Len(ss.Sorted(), 0)
}
