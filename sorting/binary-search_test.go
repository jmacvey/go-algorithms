package sorting

import (
	"testing"
)


func TestBinarySearch(t *testing.T) {
	keys := Numbers{0, 1, 2, 3, 4, 6, 8, 9}

	for i, v := range keys {
		if keys.BinarySearch(v) != i {
			t.Fatalf("BinarySearch(%v) yielded incorrect result. Expected: %v; Found: %v", v, i, keys.BinarySearch(v))
		}
	}
}