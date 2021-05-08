package sorting

func QuickSort(A []int) {
	quickSort(A, 0, len(A))
}

func quickSort(A []int, p int, r int) {
	// tail-recursive sort
	for p < r {
		q := partition(A, p, r)
		// sort left sub array
		quickSort(A, p, q)
		p = q + 1
	}
}

// do the work
// 8, 3, 5, 6, 7
func partition(A []int, p int, q int) int {
	i := p - 1 // A[i] = 8
	x := A[q-1]

	// break into [p...i] <= x; i + 1 = x; [i + 2...q) > x
	for j := p; j < q-1; j++ {
		if A[j] <= x {
			// partition has to increase to j since A[i] > A[j]
			// increment interval for i
			i++
			swap(A, j, i)
		}
	}

	// swap the last piece
	swap(A, i+1, q-1)
	return i + 1
}

func swap(A []int, i int, j int) {
	tmp := A[i]
	A[i] = A[j]
	A[j] = tmp
}
