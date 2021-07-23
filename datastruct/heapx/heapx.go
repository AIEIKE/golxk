package heapx

import "container/heap"

type HeapInt struct {
	heap []int
}

var _ heap.Interface = &HeapInt{}

func (h *HeapInt) Len() int {
	return len(h.heap)
}

func (h *HeapInt) Less(i, j int) bool {
	return h.heap[i] < h.heap[j]
}

func (h *HeapInt) Swap(i, j int) {
	h.heap[i], h.heap[j] = h.heap[j], h.heap[i]
}

func (h *HeapInt) Push(x interface{}) {
	h.heap = append(h.heap, x.(int))
}

func (h *HeapInt) Pop() interface{} {
	length := len(h.heap)
	x := h.heap[length-1]
	h.heap = h.heap[:length-2]
	return x
}
