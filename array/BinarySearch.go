package array

func BinarySearch(haystack []int, needle int) bool {
	low, middle := 0, 0
	high := len(haystack) - 1

	for low < high {
		middle = (low + high) / 2
		if haystack[middle] == needle {
			return true
		} else if haystack[middle] < needle {
			low = middle + 1
		} else {
			high = middle - 1
		}
	}
	return false
}
