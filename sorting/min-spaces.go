package sorting

func MinSpaces(pi string, favorites []string) int {
	minSpaces, ableToMatch := getMinSpaces(pi, favorites, 0)

	if ableToMatch {
		return minSpaces
	} else {
		return -1
	}
}

func getMinSpaces(pi string, favorites []string, start int) (int, bool) {
	minSpaces, ableToMatch := 0, true
	for i := start + 1; i < len(pi); i++ {
		subpi := pi[start:i]

		// the substring is in the favorites
		// so try to add a space and iterate onward
		if inFavorites(subpi, favorites) {
			// check if going forward will result in a number of min spaces
			next, ok := getMinSpaces(pi, favorites, i+1)

			if ok && next+1 < minSpaces {
				minSpaces = next + 1
			}

			if !ableToMatch {
				ableToMatch = ok
			}
		}
	}

	return minSpaces, ableToMatch
}

func inFavorites(subpi string, favorites []string) bool {
	done, index := false, 0
	for !done && index < len(favorites) {
		done = favorites[index] == subpi
	}
	return done
}
