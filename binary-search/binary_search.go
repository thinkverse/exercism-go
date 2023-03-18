package binarysearch

func SearchInts(list []int, key int) int {
	for min, max := 0, len(list); min < max; {
		guess := (min + max) / 2

		if list[guess] == key {
			return guess
		} else if list[guess] < key {
			min = guess + 1
		} else {
			max = guess
		}
	}

	return -1
}
