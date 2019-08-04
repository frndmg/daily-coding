package p5

import (
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc   string
		values [2]Any
	}{
		{
			desc:   "tuple of ints",
			values: [2]Any{1, 2},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			a, b := tc.values[0], tc.values[1]
			p := Cons(a, b)
			if !reflect.DeepEqual(a, Car(p)) || !reflect.DeepEqual(b, Cdr(p)) {
				t.Fail()
			}
		})
	}
}
