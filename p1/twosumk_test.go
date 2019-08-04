package p1

import "testing"

func Test(t *testing.T) {
	type input struct {
		list []int
		k    int
	}
	testCases := []struct {
		desc   string
		input  input
		output bool
	}{
		{
			desc: "normal example",
			input: input{
				list: []int{10, 15, 3, 7},
				k:    17,
			},
			output: true,
		},
		{
			desc: "empty",
			input: input{
				list: []int{},
				k:    0,
			},
			output: false,
		},
		{
			desc: "all one k=3",
			input: input{
				list: []int{1, 1, 1, 1, 1, 1},
				k:    3,
			},
			output: false,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			output := twosumk(tC.input.list, tC.input.k)
			if output != tC.output {
				t.Fatalf("expected %v observed %v", tC.output, output)
			}
		})
	}
}
