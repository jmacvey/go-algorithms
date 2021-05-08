package heap

// N*LOG(N)
func HeapSort(arr []int) []int {
	heap := BuildMaxHeap(arr)
	for heap.heapSize > 0 {
		heap.swap(0, heap.heapSize-1)
		heap.heapSize--
		heap.maxHeapify(0)
	}

	return heap.data[: len(arr)]
}