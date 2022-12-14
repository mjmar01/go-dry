package slices

import (
	"github.com/mjmar01/go-dry/pkg/slices"
	. "github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestRemove(t *testing.T) {
	assert := New(t)
	slice := []int{1, 2, 3, 2, 4}
	slice = slices.Remove(slice, 2)
	sort.Ints(slice)
	assert.Equal([]int{1, 3, 4}, slice)
}

func TestContains(t *testing.T) {
	assert := New(t)
	slice := []int{1, 2, 3, 4, 5, 6}
	assert.Equal(true, slices.ContainsAny(slice, 3, 7))
	assert.Equal(false, slices.ContainsAny(slice, 7, 8))
}

func TestContainsAll(t *testing.T) {
	assert := New(t)
	slice := []int{1, 2, 3, 4, 5, 6}
	assert.Equal(true, slices.ContainsAll(slice, 3, 5, 1))
	assert.Equal(false, slices.ContainsAll(slice, 3, 5, 7))
}

func TestDistinct(t *testing.T) {
	assert := New(t)
	slice := []int{1, 2, 2, 3, 3, 3}
	distinct := []int{1, 2, 3}
	assert.Equal(true, slices.ContainsAll(slices.Distinct(slice), distinct...))
	assert.Equal(3, len(slices.Distinct(slice)))
}

func TestRepeat(t *testing.T) {
	assert := New(t)
	slice := slices.Repeat(3, "1", "2")

	assert.Equal(6, len(slice))
	assert.Equal([]string{"1", "2", "1", "2", "1", "2"}, slice)
}

func TestIndexOf(t *testing.T) {
	assert := New(t)
	slice := []int{4, 2, 8, 9}

	assert.Equal(2, slices.IndexOf(slice, 8))
	assert.Equal(-1, slices.IndexOf(slice, 3))
}
