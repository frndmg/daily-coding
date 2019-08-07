package p13

import (
	"reflect"
	"testing"
)

type input struct {
	k int
	s string
}

var testCases = []struct {
	desc string
	in   input
	out  int
}{
	{
		desc: "example",
		in: input{
			k: 2,
			s: "abcba",
		},
		out: 3,
	},
	{
		desc: "all different k=2",
		in: input{
			k: 2,
			s: "abcdefghijk",
		},
		out: 2,
	},
	{
		desc: "all different k=3",
		in: input{
			k: 3,
			s: "abcdefghijk",
		},
		out: 3,
	},
	{
		desc: "k=0",
		in: input{
			k: 0,
			s: "asdfhdfjdafa",
		},
		out: 0,
	},
	{
		desc: "",
		in: input{
			k: 2,
			s: "abcddeeefghijk",
		},
		out: 5,
	},
	{
		desc: "",
		in: input{
			k: 3,
			s: "abcdeffffggggghhhhhijjjjjjkkkkkkkkkllmm",
		},
		out: 17,
	},
}

func testLongestSubstringFunc(f func(int, string) int, t *testing.T) {
	for _, tC := range testCases {
		out := f(tC.in.k, tC.in.s)
		if !reflect.DeepEqual(tC.out, out) {
			t.Fatalf("expected %v, observed %v", tC.out, out)
		}
	}
}

func TestLongestSubstringSlow(t *testing.T) {
	testLongestSubstringFunc(LongestSubstringSlow, t)
}

func TestLongestSubstring(t *testing.T) {
	testLongestSubstringFunc(LongestSubstring, t)
}
