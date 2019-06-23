package search

func BinarySearch(arr []int, value int) int {
	if len(arr) > 0 && (value < arr[0] || value > arr[len(arr)-1]) {
		return -1
	}

	lo := 0
	hi := len(arr)

	for lo < hi {
		i := ((hi - lo) / 2) + lo
		v := arr[i]

		switch {
		case value > v:
			lo = i
		case value < v:
			hi = i
		default: // value == v
			return i
		}
	}

	return -1
}
