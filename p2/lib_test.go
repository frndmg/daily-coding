package p2

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		in   []int
		out  []int
	}{
		{
			desc: "5!",
			in:   []int{1, 2, 3, 4, 5},
			out:  []int{120, 60, 40, 30, 24},
		},
		{
			desc: "only two elements",
			in:   []int{5, 6},
			out:  []int{6, 5},
		},
		{
			desc: "3, 2, 1",
			in:   []int{3, 2, 1},
			out:  []int{2, 3, 6},
		},
		{
			desc: "with cero",
			in:   []int{1, 2, 3, 0, 4},
			out:  []int{0, 0, 0, 24, 0},
		},
		{
			desc: "two ceros",
			in:   []int{0, 0, 1, 2, 3, 4},
			out:  []int{0, 0, 0, 0, 0, 0},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			out, _ := MultExcepti(tC.in)
			if !reflect.DeepEqual(out, tC.out) {
				t.Errorf("invalid answer, observed %v expected %v", out, tC.out)
			}
		})
	}
}
