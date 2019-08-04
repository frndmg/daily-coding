package p8

import (
	"reflect"
	"testing"
)

func TestDeepEqualNil(t *testing.T) {
	if !reflect.DeepEqual(nil, nil) {
		t.Fail()
	}
}

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		in   *Node
		out  int
	}{
		{
			desc: "example",
			in: &Node{
				Val:  0,
				Left: &Node{Val: 0},
				Right: &Node{
					Val: 0,
					Left: &Node{
						Val:   1,
						Left:  &Node{Val: 1},
						Right: &Node{Val: 1},
					},
					Right: &Node{Val: 0},
				},
			},
			out: 5,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			out := CountUniversalTree(tC.in)
			if !reflect.DeepEqual(tC.out, out) {
				t.Fatalf("extected %v, observed %v\n", tC.out, out)
			}
		})
	}
}
