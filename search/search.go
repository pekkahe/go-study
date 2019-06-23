package search

func BinarySearch(arr []int, value int) int {
	lo := 0
	hi := len(arr)

	for lo < hi {
		i := (hi + lo) / 2
		v := arr[i]

		switch {
		case value > v:
			lo = i + 1
		case value < v:
			hi = i
		default: // value == v
			return i
		}
	}

	return -1
}
