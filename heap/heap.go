package heap

type Heap struct {
	data     []int
	heapSize int
}

// O(n) -- tight upper bound if considering
// maxHeapify =
func BuildMaxHeap(arr []int) Heap {
	heap := Heap{
		data:     make([]int, len(arr)*2, cap(arr)*2),
		heapSize: len(arr),
	}

	for index, value := range arr {
		heap.data[index] = value
	}

	for index := len(arr) / 2; index > 0; index-- {
		heap.maxHeapify(index)
	}

	return heap
}

// LOG(N) time
func (heap *Heap) Insert(value int) {
	// insert the data at the last node in the heap
	heap.data[heap.heapSize] = value

	i := heap.heapSize
	p := heap.parent(i)
	// while the parent is < data[i]
	for p >= 0 && heap.data[p] <= value {
		// if the parent is < data[i], we are violating heap property
		// and need to move the value up a level
		heap.swap(i, p)
		i = p
		p = heap.parent(p)
	}

	heap.heapSize++
}

// LOG(N) time
func (heap *Heap) Remove(index int) {
	if index > heap.heapSize {
		return
	}

	heap.data[index] = heap.data[heap.heapSize-1]
	heap.maxHeapify(index)
}

// LOG(N) time
func (heap *Heap) maxHeapify(index int) {
	left := heap.left(index)
	right := heap.right(index)

	largest := index
	if left < heap.heapSize && heap.data[left] > heap.data[largest] {
		largest = left
	}

	if right < heap.heapSize && heap.data[right] > heap.data[largest] {
		largest = right
	}

	if largest != index {
		heap.swap(index, largest)
		heap.maxHeapify(largest)
	}
}

func (heap *Heap) left(index int) int {
	return index<<1 + 1
}

func (heap *Heap) right(index int) int {
	return index<<1 + 2
}

func (heap *Heap) parent(index int) int {
	return (index + 1>>1) - 1
}

func (heap *Heap) swap(i int, j int) {
	tmp := heap.data[i]
	heap.data[i] = heap.data[j]
	heap.data[j] = tmp
}