package leetcode

import "testing"

func TestArmstrongNumber(t *testing.T) {
	armstrongNumbers := []int{0, 1, 153, 370, 371, 407} 

	for _, v := range armstrongNumbers {
		if !IsArmstrong(v) {
			t.Fatalf("IsArmstrong(%v) returned false. Expected true.", v)
		}
	}
}