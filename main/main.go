package main

import (
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"sort"
	"time"
	pesort "github.com/pekkahe/go-study/sort"
)

func main() {
	values := make([]int, 40000)
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

	dosortcheck(a1, sorted, pesort.QuickSort)
	dosortcheck(a2, sorted, pesort.QuickSort3Way)
	dosortcheck(a3, sorted, pesort.InsertionSort)
	dosortcheck(a4, sorted, pesort.SelectionSort)
	dosortcheck(a5, sorted, pesort.BubbleSort)
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
