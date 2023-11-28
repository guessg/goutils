package containers

import (
	"container/heap"
	"errors"
)

// Heap is a wrapper to standard built-in data structure container/heap.
type Heap[T comparable] struct {
	container []T
	cmpFunc   func(i, j T) bool
}

type heapWrapper[T comparable] Heap[T]

func (hw *heapWrapper[T]) Push(x any) {
	hw.container = append(hw.container, x.(T))
}

func (hw *heapWrapper[T]) Pop() any {
	tmp := hw.container[len(hw.container)-1]
	hw.container = hw.container[:len(hw.container)-1]
	return tmp
}

func (hw *heapWrapper[T]) Len() int {
	return len(hw.container)
}

func (hw *heapWrapper[T]) Less(i, j int) bool {
	return hw.cmpFunc(hw.container[i], hw.container[j])
}

func (hw *heapWrapper[T]) Swap(i, j int) {
	hw.container[i], hw.container[j] = hw.container[j], hw.container[i]
}

// NewHeap accepts a memory allocated slice, where the capacity of slice becomes the max size
// of the heap. The elements already stored in the given slice will become the initial elements
// in the heap.
// The second paramenter accepts a function to compare the elements in the container.
// Usage example:
// ```
//
//	cmpFunc := func(i, j *Item) {
//	  return i.Value < j.Value
//	}
//	h := NewHeap(maxSize, cmpFunc)
//	err := h.Push(&Item{Value: "1234"})
//
//	// Load all the values from heap.
//	var values []*Item
//	for !h.IsEmpty() {
//	  v, _ := h.Pop()
//	  values = append(values, v)
//	}
//	// Now the values are the values in assending order.
//
// `
func NewHeap[T comparable](heapSize int, cmpFunc func(i, j T) bool) *Heap[T] {
	h := &Heap[T]{
		container: make([]T, 0, heapSize),
		cmpFunc:   cmpFunc,
	}
	heap.Init((*heapWrapper[T])(h))
	return h
}

// Push adds an element to the heap.
func (h *Heap[T]) Push(element T) error {
	if h.IsFull() {
		return errors.New("Heap full")
	}

	heap.Push((*heapWrapper[T])(h), element)
	return nil
}

// Pop returns the top item, and removes it from the heap.
func (h *Heap[T]) Pop() (T, error) {
	if h.IsEmpty() {
		var zero T
		return zero, errors.New("Heap empty")
	}

	ret := heap.Pop((*heapWrapper[T])(h))
	return ret.(T), nil
}

// Len returns the actual number of elements in the heap.
func (h *Heap[T]) Len() int {
	return len(h.container)
}

// Cap returns the capacity of the heap.
func (h *Heap[T]) Cap() int {
	return cap(h.container)
}

// Peek returns the top item in heap.
// Equivalent to the `container[0]`
func (h *Heap[T]) Peek() (T, error) {
	if h.IsEmpty() {
		var zero T
		return zero, errors.New("Heap empty")
	}

	return h.container[0], nil
}

// IsFull returns whether the heap is full.
func (h *Heap[T]) IsFull() bool {
	return h.Len() >= h.Cap()
}

// IsEmpty returns whether the heap is empty.
func (h *Heap[T]) IsEmpty() bool {
	return h.Len() <= 0
}
