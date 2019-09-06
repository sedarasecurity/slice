package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntSlice_Add(t *testing.T) {
	assert := assert.New(t)
	ss := NewIntSlice()

	ss.Add(1, 2, 3, 1)
	assert.ElementsMatch([]int{1, 2, 3}, ss.Sorted())
}

func TestIntSlice_Remove(t *testing.T) {
	assert := assert.New(t)
	ss := NewIntSlice()

	ss.Add(1, 2, 3, 4)
	assert.ElementsMatch([]int{1, 2, 3, 4}, ss.Sorted())
	ss.Remove(4)
	assert.ElementsMatch([]int{1, 2, 3}, ss.Sorted())
	ss.Remove(1)
	assert.ElementsMatch([]int{2, 3}, ss.Sorted())
}

func TestIntSlice_Contains(t *testing.T) {
	assert := assert.New(t)
	ss := NewIntSlice()

	ss.Add(1, 2, 3)
	assert.True(ss.Contains(1))
	assert.False(ss.Contains(4))
}

func TestIntSlice_Sorted(t *testing.T) {
	assert := assert.New(t)
	ss := NewIntSlice()

	ss.Add(7, 1, 5)
	assert.ElementsMatch([]int{1, 5, 7}, ss.Sorted())
}

func TestIntSlice_Reset(t *testing.T) {
	assert := assert.New(t)
	ss := NewIntSlice()

	ss.Add(1, 2, 3, 4)
	assert.ElementsMatch([]int{1, 2, 3, 4}, ss.Sorted())

	ss.Reset()
	assert.Len(ss.Sorted(), 0)
}
