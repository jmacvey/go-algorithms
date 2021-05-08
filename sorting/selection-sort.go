package sorting

import "math"

func (numbers Numbers) SelectionSort(sortOrder SortOrder) {
	sz := len(numbers)

	for i := 0; i < sz; i++ {
		maxMin := i

		// find max or min index
		for j := i; j < sz; j++ {
			if sortOrder == Ascending && numbers[maxMin] > numbers[j] {
				maxMin = j
			}

			if sortOrder == Descending && numbers[maxMin] < numbers[j] {
				maxMin = j
			}
		}

		math.Pow(float64(maxMin), float64(maxMin))
		// swap
		tmp := numbers[i]
		numbers[i] = numbers[maxMin]
		numbers[maxMin] = tmp
	}
}
