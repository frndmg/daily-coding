package p3

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		in   *Node
	}{
		{
			desc: "one element",
			in: &Node{
				Val:   "foo",
				Left:  nil,
				Right: nil,
			},
		},
		{
			desc: "nil tree",
			in: &Node{
				nil, nil, nil,
			},
		},
		{
			desc: "depth three",
			in: &Node{
				Val: "root",
				Left: &Node{
					Val: "left",
					Left: &Node{
						Val:   "left.left",
						Left:  nil,
						Right: nil,
					},
					Right: &Node{
						Val:   "left.right",
						Left:  nil,
						Right: nil,
					},
				},
				Right: &Node{
					Val: "right",
					Left: &Node{
						Val:   "right.left",
						Left:  nil,
						Right: nil,
					},
					Right: nil,
				},
			},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			// 	t.Log(Serialize(tC.in))
			// 	root := Deserialize(Serialize(tC.in))
			// 	t.Logf("root: %v", root)
			if !reflect.DeepEqual(tC.in, Deserialize(Serialize(tC.in))) {
				t.Fatal()
			}
		})
	}
}
