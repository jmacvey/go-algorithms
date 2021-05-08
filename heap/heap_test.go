package heap

import (
	"reflect"
	"testing"
)


func TestHeap(t *testing.T) { 
	arr := []int{16, 12, 14, 5, 4, 2, 1}
	expected := []int{1, 2,4,5,12,14,16}
	result := HeapSort(arr)
	if (!reflect.DeepEqual(result, expected)) {
		t.Fatalf("Heap sort did not sort the array!  Expected: %v.  Got; %v\n", expected, result)
	}

	heap := BuildMaxHeap(arr)
	heap.Insert(17)
	heap.Insert(14)
	heap.Insert(14)
	heap.Remove(4)

	expected = []int{1, 1, 2, 4, 5, 14, 14, 14, 16, 17}
	result = HeapSort(heap.data[:heap.heapSize])
	if (!reflect.DeepEqual(result, expected)) {
		t.Fatalf("Heap sort did not sort the array!  Expected: %v.  Got; %v\n", expected, result)
	}
}