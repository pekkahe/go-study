package sort

import (
	"fmt"
	"math/rand"
)

type DebugPrint struct {
	v []int
}

type DebugIterPrint struct {
	v      []int
	lo, hi int
}

type PivotSelection int

const (
	PivotFirst PivotSelection = iota
	PivotLast
	PivotMiddle
	PivotRandom
)

func QuickSort3Way(arr []int, ps PivotSelection) {
	quickSortImpl3Way(arr, 0, len(arr)-1, ps)
}

func quickSortImpl3Way(arr []int, lo, hi int, ps PivotSelection) {
	if lo < hi {
		pivot := pivot(lo, hi, ps)

		debug := DebugPrint{arr}
		debug.printArray(lo, hi, pivot)

		pivot = partition3way(arr, lo, hi, pivot)

		debug.printPartitions(lo, hi, pivot)
		debug.printLeftSide(lo, pivot)

		quickSortImpl3Way(arr, lo, pivot-1, ps)

		debug.printRightSide(pivot, hi)

		quickSortImpl3Way(arr, pivot+1, hi, ps)
	}
}

func pivot(lo, hi int, ps PivotSelection) int {
	switch ps {
	case PivotFirst:
		return lo
	case PivotLast:
		return hi
	case PivotMiddle:
		return ((hi - lo) / 2) + lo
	case PivotRandom:
		return rand.Intn(hi-lo) + lo
		// if lo%2 == 0 {
		// 	return lo
		// } else {
		// 	return hi
		// }
	default:
		panic(fmt.Sprintf("Unknown PivotSelection %d", ps))
	}
}

func partition3way(arr []int, lo, hi, p int) int {
	debug := DebugIterPrint{arr, lo, hi}

	pivot := arr[p]
	for mid := lo; mid <= hi; {
		debug.printBeginIter(lo, mid, hi)

		switch {
		case arr[mid] < pivot:
			arr[mid], arr[lo] = arr[lo], arr[mid]
			mid++
			lo++

			debug.printLess(pivot)
		case arr[mid] == pivot:
			mid++

			debug.printEqual(pivot)
		default: // arr[mid] > pivot
			arr[mid], arr[hi] = arr[hi], arr[mid]
			hi--

			debug.printGreater(pivot)
		}
	}

	return lo
}
