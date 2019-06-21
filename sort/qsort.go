package sort

func QuickSort(arr []int) {
	quickSortImpl(arr, 0, len(arr)-1)
}

func quickSortImpl(arr []int, lo, hi int) {
	if lo < hi {
		pivot := partition(arr, lo, hi)
		quickSortImpl(arr, lo, pivot-1)
		quickSortImpl(arr, pivot+1, hi)
	}
}

func partition(arr []int, lo, hi int) int {
	// Partitioning assumes that pivot is always the last element
	pivot := hi

	for pivot > lo {
		if arr[lo] < arr[pivot] {
			lo++
		} else {
			// Swap pivot and value
			arr[pivot], arr[lo] = arr[lo], arr[pivot]
			// Decrement pivot
			pivot--
			// Move pivot back before value
			arr[lo], arr[pivot] = arr[pivot], arr[lo]
		}
	}

	return pivot
}

