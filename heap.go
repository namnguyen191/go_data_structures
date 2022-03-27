package main

import "errors"

type MinIntHeap interface {
	Peek() (int, error)
	Poll() (int, error)
	Add(item int)
}

type minIntHeap struct {
	items []int
}

func NewMinIntHeap(numbers []int) MinIntHeap {
	heap := &minIntHeap{}

	for _, num := range numbers {
		heap.Add(num)
	}

	return heap
}

func (*minIntHeap) getLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (*minIntHeap) getRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (*minIntHeap) getParentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h *minIntHeap) getLeftChildValue(index int) int {
	return h.items[h.getLeftChildIndex(index)]
}

func (h *minIntHeap) getRightChildValue(index int) int {
	return h.items[h.getRightChildIndex(index)]
}

func (h *minIntHeap) getParentValue(index int) int {
	return h.items[h.getParentIndex(index)]
}

func (h *minIntHeap) hasLeftChild(parentIndex int) bool {
	return h.getLeftChildIndex(parentIndex) < len(h.items)
}

func (h *minIntHeap) hasRightChild(parentIndex int) bool {
	return h.getRightChildIndex(parentIndex) < len(h.items)
}

func (h *minIntHeap) hasParent(childIndex int) bool {
	return h.getParentIndex(childIndex) >= 0
}

func (h *minIntHeap) swap(indexOne int, indexTwo int) {
	temp := h.items[indexOne]
	h.items[indexOne] = h.items[indexTwo]
	h.items[indexTwo] = temp
}

func (h *minIntHeap) Peek() (int, error) {
	if len(h.items) == 0 {
		return -1, errors.New("heap is empty")
	}

	return h.items[0], nil
}

func (h *minIntHeap) heapifyDown() {
	index := 0

	// check left child first because if there is no left child
	// then there is no right child
	for h.hasLeftChild(index) {
		smallerChildIndex := h.getLeftChildIndex(index)

		if h.hasRightChild(index) &&
			h.getRightChildValue(index) < h.getLeftChildValue(index) {
			smallerChildIndex = h.getRightChildIndex(index)
		}

		if h.items[index] < h.items[smallerChildIndex] {
			break
		} else {
			h.swap(index, smallerChildIndex)
			index = smallerChildIndex
		}
	}
}

func (h *minIntHeap) heapifyUp() {
	index := len(h.items) - 1

	for h.hasParent(index) && h.getParentValue(index) > h.items[index] {
		h.swap(h.getParentIndex(index), index)
		index = h.getParentIndex(index)
	}
}

func (h *minIntHeap) Poll() (int, error) {
	size := len(h.items)

	if size == 0 {
		return -1, errors.New("heap is empty")
	}

	item := h.items[0]

	h.items[0] = h.items[size-1]

	h.items = h.items[:len(h.items)-1]

	h.heapifyDown()

	return item, nil
}

func (h *minIntHeap) Add(item int) {
	h.items = append(h.items, item)

	h.heapifyUp()
}
