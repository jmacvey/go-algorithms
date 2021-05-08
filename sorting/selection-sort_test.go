package sorting

import (
	"reflect"
	"testing"
)

func TestAscendingSelectionSort(t *testing.T) {
	keys := Numbers{0, 1, 3, 4, 1, 2, 6, 8, 9}
	want := Numbers{0, 1, 1, 2, 3, 4, 6, 8, 9}
	keys.SelectionSort(Ascending)
	if !reflect.DeepEqual(keys, want) {
		t.Fatalf(`Arrays do not match. Found: %v; Wanted: %v`,keys, want)
	}
}

func TestDescendingSelectionSort(t *testing.T) {
	keys := Numbers{0, 1, 3, 4, 1, 2, 6, 8, 9}
	want := Numbers{9, 8, 6, 4, 3, 2, 1, 1, 0}
	keys.SelectionSort(Descending)
	if !reflect.DeepEqual(keys, want) {
		t.Fatalf(`Arrays do not match. Found: %v; Wanted: %v`,keys, want)
	}
}
