package sorting

import (
	"reflect"
	"testing"
)

func TestPartition(t *testing.T) {
	keys := Numbers{14, 28, 7, 42, 12, 16}
	expected := Numbers{14, 7, 12, 16, 28, 42}
	partition(keys, 0, len(keys))
	if !reflect.DeepEqual(keys, expected) {
		t.Fatalf(`Arrays do not match. Found %v; Wanted: %v`, keys, expected)
	}
}

func TestQuickSort(t *testing.T) {
	// 1, 6, 7, 9, 4, 5
	// V := 1
	arr := []int{1, 6, 7, 9, 4, 5}
	expected := []int{1, 4, 5, 6, 7, 9}
	QuickSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Fatalf(`Arrays do not match. Found %v; Wanted: %v`, arr, expected)
	}

	arr = []int{2,7,4,0,7,4,2,7,2,3}
	expected = []int{0,2,2,2,3,4,4,7,7,7}
	QuickSort(arr)

	if !reflect.DeepEqual(arr, expected) {
		t.Fatalf(`Arrays do not match. Found %v; Wanted: %v`, arr, expected)
	}


}