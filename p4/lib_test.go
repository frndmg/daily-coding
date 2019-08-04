package p4

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		in   []int
		out  int
	}{
		{
			desc: "example 1",
			in:   []int{3, 4, -1, 1},
			out:  2,
		},
		{
			desc: "example 2",
			in:   []int{1, 2, 0},
			out:  3,
		},
		{
			desc: "empty",
			in:   []int{},
			out:  1,
		},
		{
			desc: "one element",
			in:   []int{0},
			out:  1,
		},
		{
			desc: "3, 4, 2, -1, 1",
			in:   []int{3, 4, 2, -1, 1},
			out:  5,
		},
		{
			desc: "5, 4, 3, 2, 1",
			in:   []int{5, 4, 3, 2, 1},
			out:  6,
		},
		{
			desc: "6, 5, 4, 1, 3, 2",
			in:   []int{6, 5, 4, 1, 3, 2},
			out:  7,
		},
		{
			desc: "foo",
			in:   []int{6, 1, 2, 3, 4, 5},
			out:  7,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			out := FindMissingMinPositiveV2(tC.in)
			if !reflect.DeepEqual(out, tC.out) {
				t.Fatalf("Unexpected %v, expected %v\n", out, tC.out)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	n := 10000
	length := 1000
	min := -5000
	max := 5000

	rand.Seed(42)

	for i := 0; i < n; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			l := make([]int, length)
			for i := 0; i < length; i++ {
				l[i] = randomInt(min, max)
			}

			out1 := FindMissingMinPositiveV2(copyL(l))
			out2 := SolveExtraMemory(copyL(l))

			if out1 != out2 {
				t.Fatalf("Results mismatch %v != %v\n", out1, out2)
			}
		})
	}
}

func randomInt(a, b int) int {
	return rand.Intn(b-a) + a
}

func copyL(l []int) []int {
	l2 := make([]int, len(l))
	copy(l2, l)
	return l2
}
