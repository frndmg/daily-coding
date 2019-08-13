package nhouses

import "testing"

type input struct {
	cost [][]int
	N    int
	K    int
}

var testCases = []struct {
	desc string
	in   input
	out  int
}{
	{
		desc: "example",
		in: input{
			cost: [][]int{
				{1, 2, 1},
				{2, 3, 4},
			},
			N: 2,
			K: 3,
		},
		out: 3,
	},
	{
		desc: "example 2",
		in: input{
			cost: [][]int{
				{1, 2, 3, 4, 5},
				{1, 4, 3, 2, 3},
				{4, 4, 1, 1, 1},
			},
			N: 3,
			K: 5,
		},
		out: 4,
	},
}

func MakeFromSlice(ss ...[]int) [][]int {
	return ss
}

func testNHousesFunc(f func([][]int, int, int) int, t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			out := f(tC.in.cost, tC.in.N, tC.in.K)
			if tC.out != out {
				t.Errorf("extected %v, observed %v", tC.out, out)
			}
		})
	}
}

func TestBacktrack(t *testing.T) {
	testNHousesFunc(MinNHousesBacktrack, t)
}

func TestDisktra(t *testing.T) {
	testNHousesFunc(MinNHousesDisktra, t)
}

func TestDP(t *testing.T) {
	testNHousesFunc(MinNHousesDP, t)
}
