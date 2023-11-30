package containers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainWithExistingElement(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	assert.True(t, Contain(slice, 3))
}

func TestContainWithNonExistingElement(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	assert.False(t, Contain(slice, 6))
}

func TestFindWithExistingElement(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	assert.Equal(t, 2, Find(slice, 3))
}

func TestFindWithNonExistingElement(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	assert.Equal(t, -1, Find(slice, 6))
}

func TestFirstWithNonEmptySlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	assert.Equal(t, 1, First(slice))
}

func TestLastWithNonEmptySlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	assert.Equal(t, 5, Last(slice))
}

func TestLastWithEmptySlice(t *testing.T) {
	var slice []int
	assert.Equal(t, 0, Last[int](slice))
}

func TestRepeatWithPositiveCount(t *testing.T) {
	assert.Equal(t, []int{2, 2, 2}, Repeat(3, 2))
}

func TestRepeatWithZeroCount(t *testing.T) {
	assert.Nil(t, Repeat(0, 2))
}

func TestInterfaceSlice(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	expected := []interface{}{1, 2, 3, 4, 5}
	assert.Equal(t, expected, InterfaceSlice(slice))
}

func TestClone(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	cloned := Clone(slice)
	assert.Equal(t, slice, cloned)
	assert.NotSame(t, slice, cloned)
}

func TestReverse(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	reversed := Reverse(slice)
	assert.Equal(t, []int{5, 4, 3, 2, 1}, reversed)
}
