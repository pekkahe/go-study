package search_test

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"testing"
	"time"

	. "github.com/pekkahe/go-study/search"
	. "github.com/pekkahe/go-study/testutil"
)

var sorted []int
var value, expected int

func TestMain(m *testing.M) {
	n := 5000000
	v := make([]int, n)
	for i := 0; i < len(v); i++ {
		v[i] = i + 1 - (n / 2)
	}
	sorted = v

	rand.Seed(time.Now().UnixNano())
	value = rand.Int()%n - (n / 2)

	start := time.Now()
	expected = sort.SearchInts(sorted, value)
	elapsed := time.Since(start)

	fmt.Printf("Reference: %s\n", elapsed)

	os.Exit(m.Run())
}

func TestBinarySearch(t *testing.T) {
	start := time.Now()
	v := BinarySearch(sorted, value)
	elapsed := time.Since(start)

	fmt.Printf("%s: %s\n", FunctionName(BinarySearch), elapsed)

	AssertVEqual(t, expected, v)
}
