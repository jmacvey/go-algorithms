package sorting

import (
	"testing"
)


func TestSumTwo(t *testing.T) {
	keys := Numbers{0, 1, 3, 4, 1, 2, 6, 8, 9}

	a, b := keys.Sum2(3)

	if (a != 0 || b != 4) {
		t.Fatalf("Invalid indices found. Expected (%v, %v); Found: (%v, %v).", 0, 4, a, b)
	}
}