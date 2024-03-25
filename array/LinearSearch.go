package array

func LinearSearch(haystack []int, needle int) bool {
	for e := range haystack {
		if needle == e {
			return true
		}
	}
	return false
}
