package sort

func QuickSort3Way(arr []int) {
	quickSortImpl(arr, 0, len(arr)-1)
}

func quickSortImpl3Way(arr []int, lo, hi int) {
	if lo < hi {
		pivot := partition3way(arr, lo, hi, hi)

		//padding := lo * 2
		// fmt.Printf("Current:     %v\n", arr)
		// fmt.Printf("Partition:   %s%v pivot: [%d]\n",
		// 	strings.Repeat(" ", padding),
		// 	arr[lo:hi+1],
		// 	pivot)
		// fmt.Printf("Left/Right:  %s%v %v\n",
		// 	strings.Repeat(" ", padding),
		// 	arr[lo:pivot],
		// 	arr[pivot+1:hi+1])

		// fmt.Printf("Left:        %s%v ->\n",
		// 	strings.Repeat(" ", padding),
		// 	arr[lo:pivot])

		quickSortImpl3Way(arr, lo, pivot-1)

		// fmt.Printf("Right:       %s%v ->\n",
		// 	strings.Repeat(" ", (pivot+1)*2),
		// 	arr[pivot+1:hi+1])

		quickSortImpl3Way(arr, pivot+1, hi)
	}
}

func partition3way(arr []int, lo, hi, p int) int {
	// padding := lo * 2
	// t1, t2 := lo, hi
	// fmt.Printf("  Partition: %s%v indices: [%d-%d] pivot: [%d]=%d\n",
	// 	strings.Repeat(" ", padding),
	// 	arr[lo:hi+1],
	// 	lo,
	// 	hi,
	// 	p,
	// 	arr[p])

	pivot := arr[p]
	for mid := lo; mid <= hi; {

		// marks := make([]string, len(arr))
		// for i := range marks {
		// 	marks[i] = " "
		// }
		// marks[lo] = "L"
		// marks[mid] = "M"
		// marks[hi] = "H"
		// fmt.Printf("             %s%v\n",
		// 	strings.Repeat(" ", padding),
		// 	arr[t1:t2+1])
		// fmt.Printf("              %s%s",
		// 	strings.Repeat(" ", padding),
		// 	strings.Join(marks[t1:t2+1], " "))

		switch {
		case arr[mid] < pivot:
			arr[mid], arr[lo] = arr[lo], arr[mid]
			mid++
			lo++

			//fmt.Printf("    M < %d  -> swap(M,L), M++, L++", pivot)
		case arr[mid] == pivot:
			mid++

			//fmt.Printf("    M == %d -> M++", pivot)
		default: // arr[mid] > pivot
			arr[mid], arr[hi] = arr[hi], arr[mid]
			hi--

			///fmt.Printf("    M > %d  -> swap(M,H), H--", pivot)
		}
		//fmt.Print("\n")
	}
	// fmt.Printf("  Done:      %s%v\n",
	// 	strings.Repeat(" ", padding),
	// 	arr[t1:t2+1])

	return lo
}
