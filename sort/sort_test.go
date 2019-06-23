package sort_test

import (
	"fmt"
	"os"
	"sort"
	"testing"
	"time"

	. "github.com/pekkahe/go-study/sort"
	. "github.com/pekkahe/go-study/testutil"
)

var unsorted []int
var sorted []int

func TestMain(m *testing.M) {
	unsorted = RandomNumbers(50000)
	sorted = append([]int(nil), unsorted...)

	start := time.Now()
	sort.Ints(sorted)
	elapsed := time.Since(start)

	fmt.Printf("Reference: %s\n", elapsed)

	os.Exit(m.Run())
}

func TestQuickSort(t *testing.T) {
	runTest(t, QuickSort)
}

func TestQuickSort3Way(t *testing.T) {
	fmt.Print("QuickSort3Way.PivotFirst.")
	runTest(t, func(v []int) { QuickSort3Way(v, PivotFirst) })

	fmt.Print("QuickSort3Way.PivotLast.")
	runTest(t, func(v []int) { QuickSort3Way(v, PivotLast) })

	fmt.Print("QuickSort3Way.PivotMiddle.")
	runTest(t, func(v []int) { QuickSort3Way(v, PivotMiddle) })

	fmt.Print("QuickSort3Way.PivotRandom.")
	runTest(t, func(v []int) { QuickSort3Way(v, PivotRandom) })
}

func TestInsertionSort(t *testing.T) {
	runTest(t, InsertionSort)
}

func TestSelectionSort(t *testing.T) {
	runTest(t, SelectionSort)
}

func TestBubbleSort(t *testing.T) {
	runTest(t, BubbleSort)
}

func runTest(t *testing.T, sortF func([]int)) {
	v := append([]int(nil), unsorted...)

	start := time.Now()
	sortF(v)
	elapsed := time.Since(start)

	fmt.Printf("%s: %s\n", FunctionName(sortF), elapsed)

	AssertEqual(t, sorted, v)
}
