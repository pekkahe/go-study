package testutil

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"
)

func RandomNumbers(n int) []int {
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

func FunctionName(i interface{}) string {
	fullName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	parts := strings.Split(fullName, ".")
	return parts[len(parts)-1]
}

func AssertEqual(t *testing.T, expected, actual []int) {
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

func AssertVEqual(t *testing.T, expected, actual interface{}) {
	if actual != expected {
		t.Errorf("actual != expected => %d != %d\n", actual, expected)
		return
	}
}
