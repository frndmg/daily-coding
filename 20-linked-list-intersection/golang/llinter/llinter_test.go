package llinter

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	desc string
	a    *LList
	b    *LList
	out  *LList
}{
	{
		desc: "example",
		a:    MakeLList(1, 2, 3, 4),
		b:    MakeLList(5, 6, 1, 3, 4),
		out:  MakeLList(3, 4),
	},
	{
		desc: "example 2",
		a:    MakeLList(3, 5, 6, 4, 2, 7),
		b:    MakeLList(1, 9, 3, 4, 2, 7),
		out:  MakeLList(4, 2, 7),
	},
	{
		desc: "example 3",
		a:    MakeLList(3, 5, 6, 4, 2, 1),
		b:    MakeLList(9, 8, 7, 5, 6, 4, 2, 1),
		out:  MakeLList(5, 6, 4, 2, 1),
	},
	{
		desc: "empty result",
		a:    MakeLList(1),
		b:    MakeLList(2, 3),
		out:  nil,
	},
}

func MakeLList(s ...int) *LList {
	if len(s) == 0 {
		return nil
	}

	r := &LList{s[0], nil}
	n := r
	for i := 1; i < len(s); i++ {
		n.next = &LList{s[i], nil}
		n = n.next
	}
	return r
}

func Test(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			out := Intersection(tC.a, tC.b)
			if !reflect.DeepEqual(tC.out, out) {
				t.Errorf("expected %v, observed %v\n", tC.out, out)
			}
		})
	}
}
