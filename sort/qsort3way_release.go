// +build !debug

package sort

func (d *DebugPrint) printArray(lo, hi, pivot int) {
}

func (d *DebugPrint) printPartitions(lo, hi, pivot int) {
}

func (d *DebugPrint) printLeftSide(lo, pivot int) {
}

func (d *DebugPrint) printRightSide(pivot, hi int) {
}

func (d *DebugIterPrint) printBeginIter(lo, mid, hi int) {
}

func (d *DebugIterPrint) printLess(pivot int) {
}

func (d *DebugIterPrint) printEqual(pivot int) {
}

func (d *DebugIterPrint) printGreater(pivot int) {
}
