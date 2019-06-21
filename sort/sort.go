package main

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"time"
)

func main() {
	values := make([]int, 10000)
	for i := 0; i < len(values); i++ {
		values[i] = rand.Int() % 10
	}
	//values = []int{0, 1, 1, 1, 7, 0, 8, 8, 6, 9, 4, 7, 4, 6, 5, 3, 8, 1, 0, 1}
	//values = []int{7, 9, 7, 8}

	const maxPrint int = 20
	fmt.Printf("Values [n=%d]: ", len(values))
	if len(values) > maxPrint {
		fmt.Printf("%v ...\n", values[:maxPrint])
	} else {
		fmt.Printf("%v\n", values)
	}

	sorted := append([]int(nil), values...)
	start := time.Now()
	sort.Ints(sorted)
	elapsed := time.Since(start)
	fmt.Printf("Reference: %s\n", elapsed)

	a1 := append([]int(nil), values...)
	a2 := append([]int(nil), values...)
	a3 := append([]int(nil), values...)
	a4 := append([]int(nil), values...)
	a5 := append([]int(nil), values...)

	dosortcheck(a1, sorted, quickSort1)
	dosortcheck(a2, sorted, quickSort2)
	dosortcheck(a3, sorted, insertionSort)
	dosortcheck(a4, sorted, selectionSort)
	dosortcheck(a5, sorted, bubbleSort)
}

func quickSort1(arr []int) {
	quickSortImpl1(arr, 0, len(arr)-1)
}

func quickSortImpl1(arr []int, lo, hi int) {
	if lo < hi {
		pivot := partition1(arr, lo, hi)
		quickSortImpl1(arr, lo, pivot-1)
		quickSortImpl1(arr, pivot+1, hi)
	}
}

func partition1(arr []int, lo, hi int) int {
	pivot := hi
	//fmt.Printf("Partition: %v Pivot: [%d]=%d\n", arr[lo:hi+1], pivot, arr[pivot])

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
	//fmt.Printf("Done:      %v\n", arr[lo:hi+1])
	//fmt.Printf("Array:     %v\n", arr)

	return pivot
}

func quickSort2(arr []int) {
	quickSortImpl2(arr, 0, len(arr)-1)
}

func quickSortImpl2(arr []int, lo, hi int) {
	if lo < hi {
		pivot := partition2(arr, lo, hi, hi)

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

		quickSortImpl2(arr, lo, pivot-1)

		// fmt.Printf("Right:       %s%v ->\n",
		// 	strings.Repeat(" ", (pivot+1)*2),
		// 	arr[pivot+1:hi+1])

		quickSortImpl2(arr, pivot+1, hi)
	}
}

func partition2(arr []int, lo, hi, p int) int {
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

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = key
	}
}

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func bubbleSort(arr []int) {
	for end := len(arr); end > 0; end-- {
		for i := 1; i < end; i++ {
			if arr[i-1] > arr[i] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
	}
}

func dosortcheck(a, sorted []int, f func([]int)) {
	name := getFunctionName(f)
	elapsed := dosort(a, f)
	_, err := checkEqual(sorted, a)

	fmt.Printf("%s: %s", name, elapsed)
	if err == nil {
		fmt.Print(" -> SUCCESS")
	} else {
		fmt.Printf(" -> FAIL\n%s", err)
	}
	fmt.Println()
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func dosort(a []int, f func([]int)) time.Duration {
	start := time.Now()
	f(a)
	elapsed := time.Since(start)
	return elapsed
}

func checkEqual(expected, actual []int) (bool, error) {
	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			var b strings.Builder
			b.WriteString(fmt.Sprintf("  ERROR:    actual[%d] != expected[%d] => %d != %d\n", i, i, actual[i], expected[i]))
			//b.WriteString(fmt.Sprintf("  Expected: %v\n", expected))
			//b.WriteString(fmt.Sprintf("  Actual:   %v", actual))
			return false, errors.New(b.String())
		}
	}

	return true, nil
}
