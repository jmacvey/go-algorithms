package sorting

func (numbers Numbers) BinarySearch(value int) int {
	p, r, i := 0, len(numbers), -1

	for p != r {
		// [p, r) -> [p, h], [h, r)
		h := (p + r) / 2

		if numbers[h] < value {
			p = h // search [h, r)
		} else if numbers[h] > value {
			r = h // search [p, h)
		} else {
			i = h
			p = r
		}
	}

	return i
}