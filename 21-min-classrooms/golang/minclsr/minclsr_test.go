package minclsr

import "testing"

var testCases = []struct {
	desc string
	in   [][]int
	out  int
}{
	{
		desc: "example",
		in: [][]int{
			{30, 75},
			{0, 50},
			{60, 150},
		},
		out: 2,
	},
}

func Test(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			out := MinClassrooms(tC.in)
			if tC.out != out {
				t.Errorf("expected %v, observed %v\n", tC.out, out)
			}
		})
	}
}
