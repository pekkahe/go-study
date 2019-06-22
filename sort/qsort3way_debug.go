// +build debug

package sort

import (
	"fmt"
	"strings"
)

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
