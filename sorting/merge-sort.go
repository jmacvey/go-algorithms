package sorting

func (numbers Numbers) MergeSort() {
	sorted := numbers.MergeSortRecursive(0, len(numbers), make([]int, len(numbers)))

	for i := 0; i < len(numbers); i++ {
		numbers[i] = sorted[i]
	}
}

func (numbers Numbers) MergeSortRecursive(p int, r int, arr []int) []int {
	// base case
	if p == r {
		return arr
	}

	q := (p + r) / 2

	// sort [p, q)
	numbers.MergeSortRecursive(p, q, arr)

	// sort [q, r)
	numbers.MergeSortRecursive(q+1, r, arr)
	return numbers.combine(p, q, r, arr) // O(N)
}

func (numbers Numbers) combine(p int, q int, r int, arr []int) []int {
	a, b, i := p, q, p

	for i < r {
		if a < q && b < r {
			if numbers[a] < numbers[b] {
				arr[i] = numbers[a]
				a++
			} else {
				arr[i] = numbers[b]
				b++
			}
		} else if a < q {
			arr[i] = numbers[a]
			a++
		} else {
			arr[i] = numbers[b]
			b++
		}

		i++
	}

	return arr
}
