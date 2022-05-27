package sol

import "container/heap"

type MinIntHeap []int

func (h *MinIntHeap) Len() int {
	return len(*h)
}
func (h *MinIntHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}
func (h *MinIntHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *MinIntHeap) Push(val interface{}) {
	*h = append(*h, val.(int))
}
func (h *MinIntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[1 : n-1]
	return x
}
func findKthLargest(nums []int, k int) int {
	pq := MinIntHeap(nums[:k])
	heap.Init(&pq)
	for _, num := range nums {
		if pq[0] < num {
			pq[0] = num
			// heapify
			heap.Fix(&pq, 0)
		}
	}
	return heap.Pop(&pq).(int)
}
