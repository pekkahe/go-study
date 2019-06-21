package sort_test

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"testing"
	"time"

	pesort "github.com/pekkahe/go-study/sort"
)

var unsorted []int
var sorted []int

func TestMain(m *testing.M) {
	unsorted = generateSingleDigitNumbers(10)
	sorted = append([]int(nil), unsorted...)

	start := time.Now()
	sort.Ints(sorted)
	elapsed := time.Since(start)

	fmt.Printf("Reference: %s\n", elapsed)

	os.Exit(m.Run())
}

func generateSingleDigitNumbers(n int) []int {
	v := make([]int, n)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(v); i++ {
		v[i] = rand.Int() % 10
	}

	const printMaxN int = 20
	fmt.Printf("Numbers (n=%d): ", len(v))
	if len(v) > printMaxN {
		fmt.Printf("%v ...\n", v[:printMaxN])
	} else {
		fmt.Printf("%v\n", v)
	}

	return v
}

func TestQuickSort(t *testing.T) {
	runTest(t, pesort.QuickSort)
}

func TestQuickSort3Way(t *testing.T) {
	runTest(t, pesort.QuickSort3Way)
}

func TestInsertionSort(t *testing.T) {
	runTest(t, pesort.InsertionSort)
}

func TestSelectionSort(t *testing.T) {
	runTest(t, pesort.SelectionSort)
}

func TestBubbleSort(t *testing.T) {
	runTest(t, pesort.BubbleSort)
}

func runTest(t *testing.T, sortF func([]int)) {
	v := append([]int(nil), unsorted...)

	start := time.Now()
	sortF(v)
	elapsed := time.Since(start)

	fmt.Printf("%s: %s\n", getFunctionName(sortF), elapsed)

	assertEqual(t, sorted, v)
}

func getFunctionName(i interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	parts := strings.Split(fullName, ".")
	return parts[len(parts)-1]
}

func assertEqual(t *testing.T, expected, actual []int) {
	if len(actual) != len(expected) {
		t.Errorf("len(actual) != len(expected) => %d != %d\n",
			len(actual),
			len(expected))
		return
	}

	var errors []int
	const printErrors = 3

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			errors = append(errors, i)
			if len(errors) > printErrors {
				break
			}
		}
	}

	for i := 0; i < len(errors) && i < printErrors; i++ {
		idx := errors[i]
		t.Errorf("actual[%d] != expected[%d] => %d != %d\n", idx, idx, actual[idx], expected[idx])
	}

	if len(errors) > printErrors {
		t.Errorf("...\n")
	}
}
