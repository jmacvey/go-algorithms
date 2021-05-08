package sorting

func (keys Numbers) InsertionSort(sortOrder SortOrder) {
	for j := 1; j < len(keys); j++ {
		key := keys[j]
		i := j - 1
		for i >= 0 && ((sortOrder == Ascending && key < keys[i]) || (sortOrder == Descending && key > keys[i])) {
			keys[i+1] = keys[i]
			i--

		}
		keys[i+1] = key
	}
}
