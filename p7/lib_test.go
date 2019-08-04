package p7

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		out  int
	}{
		{
			desc: "example",
			in:   "111",
			out:  3,
		},
		{
			desc: "example",
			in:   "11111",
			out:  8,
		},
		{
			desc: "example 3",
			in:   "323",
			out:  2,
		},
		{
			desc: "",
			in:   "",
			out:  0,
		},
		{
			desc: "9",
			in:   "9",
			out:  1,
		},
		{
			desc: "",
			in:   "22322",
			out:  6,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			out := Count(tC.in)
			if !reflect.DeepEqual(tC.out, out) {
				t.Fatalf("expected %v, received %v", tC.out, out)
			}
		})
	}
}
