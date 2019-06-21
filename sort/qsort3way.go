package sort

import (
	"fmt"
	"strings"
)

type PivotSelection int

const (
	PivotStart PivotSelection = iota
	PivotMiddle
	PivotEnd
)

func QuickSort3Way(arr []int, ps PivotSelection) {
	quickSortImpl3Way(arr, 0, len(arr)-1, ps)
}

func quickSortImpl3Way(arr []int, lo, hi int, ps PivotSelection) {
	if lo < hi {
		pivot := pivot(lo, hi, ps)

		//debug := DebugPrint{arr}
		//debug.printArray(lo, hi, pivot)

		pivot = partition3way(arr, lo, hi, pivot)

		//debug.printPartitions(lo, hi, pivot)
		//debug.printLeftSide(lo, pivot)

		quickSortImpl3Way(arr, lo, pivot-1, ps)

		//debug.printRightSide(pivot, hi)

		quickSortImpl3Way(arr, pivot+1, hi, ps)
	}
}

func pivot(lo, hi int, ps PivotSelection) int {
	switch ps {
	case PivotStart:
		return lo
	case PivotMiddle:
		return ((hi - lo) / 2) + lo
	default: // PivotEnd
		return hi
	}
}

func partition3way(arr []int, lo, hi, p int) int {
	//debug := DebugIterPrint{arr, lo, hi}

	pivot := arr[p]
	for mid := lo; mid <= hi; {
		//debug.printBeginIter(lo, mid, hi)

		switch {
		case arr[mid] < pivot:
			arr[mid], arr[lo] = arr[lo], arr[mid]
			mid++
			lo++

			//debug.printLess(pivot)
		case arr[mid] == pivot:
			mid++

			//debug.printEqual(pivot)
		default: // arr[mid] > pivot
			arr[mid], arr[hi] = arr[hi], arr[mid]
			hi--

			//debug.printGreater(pivot)
		}
	}

	return lo
}

type DebugPrint struct {
	v []int
}

func (d *DebugPrint) printArray(lo, hi, pivot int) {
	fmt.Printf("Array:       %v ", d.v)
	fmt.Printf("Next: [%d-%d] Pivot: [%d]=%d\n",
		lo,
		hi,
		pivot,
		d.v[pivot])
}

func (d *DebugPrint) printPartitions(lo, hi, pivot int) {
	leftSide := d.v[lo:pivot]
	rightSide := d.v[pivot+1 : hi+1]

	margin := 1
	if len(leftSide) < 1 || len(rightSide) < 1 {
		margin = 0
	}

	fmt.Printf("Left/Right:  %s%v%s%v Pivot: [%d]=%d\n",
		strings.Repeat(" ", lo*2),
		leftSide,
		strings.Repeat(" ", margin),
		rightSide,
		pivot,
		d.v[pivot])
}

func (d *DebugPrint) printLeftSide(lo, pivot int) {
	fmt.Printf("Left:        %s%v ->\n",
		strings.Repeat(" ", lo*2),
		d.v[lo:pivot])
}

func (d *DebugPrint) printRightSide(pivot, hi int) {
	fmt.Printf("Right:       %s%v ->\n",
		strings.Repeat(" ", (pivot+1)*2),
		d.v[pivot+1:hi+1])
}

type DebugIterPrint struct {
	v      []int
	lo, hi int
}

func (d *DebugIterPrint) printBeginIter(lo, mid, hi int) {
	fmt.Printf("             %s%v\n",
		strings.Repeat(" ", d.lo*2),
		d.v[d.lo:d.hi+1])

	marks := make([]string, len(d.v))
	for i := range marks {
		marks[i] = " "
	}
	marks[lo] = "L"
	marks[mid] = "M"
	marks[hi] = "H"

	fmt.Printf("              %s%s",
		strings.Repeat(" ", d.lo*2),
		strings.Join(marks[d.lo:d.hi+1], " "))
}

func (d *DebugIterPrint) printLess(pivot int) {
	fmt.Printf("  -> M < %d  -> swap(M,L), M++, L++\n", pivot)
}

func (d *DebugIterPrint) printEqual(pivot int) {
	fmt.Printf("  -> M == %d -> M++\n", pivot)
}

func (d *DebugIterPrint) printGreater(pivot int) {
	fmt.Printf("  -> M > %d  -> swap(M,H), H--\n", pivot)
}
