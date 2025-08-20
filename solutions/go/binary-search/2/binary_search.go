package binarysearch

func SearchInts(list []int, key int) int {
	for left, right := 0, len(list); left != right; {
		mid := (left + right) / 2

		switch {
		case list[mid] == key:
			return mid
		case key < list[mid]:
			right = mid
		default:
			left = mid + 1
		}
	}
	return -1
}
