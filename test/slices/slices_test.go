package slices

import (
	"go-dry/pkg/slices"
	"testing"
)
import . "github.com/stretchr/testify/assert"

func TestElementInSlice(t *testing.T) {
	assert := New(t)
	slice := []int{1, 2, 3, 4, 5, 6}
	assert.Equal(true, slices.ElementInSlice(slice, 3))
	assert.Equal(false, slices.ElementInSlice(slice, 7))
}

func TestDistinct(t *testing.T) {
	assert := New(t)
	slice := []int{1, 2, 2, 3, 3, 3}
	distinct := []int{1, 2, 3}
	assert.Equal(distinct, slices.Distinct(slice))
}
