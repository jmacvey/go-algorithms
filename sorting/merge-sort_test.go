package sorting

import (
	"reflect"
	"testing"
)

func TestAscendingMergeSort(t *testing.T) {
	keys := Numbers{0, 1, 3, 4, 1, 2, 6, 8, 9}
	want := Numbers{0, 1, 1, 2, 3, 4, 6, 8, 9}
	keys.MergeSort()
	if !reflect.DeepEqual(keys, want) {
		t.Fatalf(`Arrays do not match. Found: %v; Wanted: %v`,keys, want)
	}
}
