package justtext

import (
	"reflect"
	"testing"
)

type input struct {
	ws []string
	k  int
}

var testCases = []struct {
	desc string
	in   input
	out  []string
}{
	{
		desc: "example",
		in: input{
			ws: []string{
				"the", "quick", "brown",
				"fox", "jumps", "over",
				"the", "lazy", "dog",
			},
			k: 16,
		},
		out: []string{
			"the  quick brown",
			"fox  jumps  over",
			"the   lazy   dog",
		},
	},
	{
		desc: "one word in line",
		in: input{
			ws: []string{
				"the",
			},
			k: 4,
		},
		out: []string{"the "},
	},
}

func Test(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			out := Justify(tC.in.ws, tC.in.k)
			if !reflect.DeepEqual(tC.out, out) {
				t.Errorf("expected %v, observed %v\n", tC.out, out)
			}
		})
	}
}
