package sorting

import (
	"fmt"
)


func (numbers Numbers) Sum2(x int) (int, int) {
	fmt.Println("Finding sums")
	numbers.MergeSort()
	i, j, p, r := -1, -1, 0, len(numbers)-1

	for p != r {
		sum := numbers[p] + numbers[r]
		if sum < x {
			p++
		} else if sum > x {
			r--
		} else {
			i, j = p, r
			p = r
		}
	}

	return i, j
}
