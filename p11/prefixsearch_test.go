package p11

import (
	"reflect"
	"sort"
	"testing"
)

type input struct {
	s string
	a []string
}

var testCases = []struct {
	desc string
	in   input
	out  []string
}{
	{
		desc: "example",
		in: input{
			s: "de",
			a: []string{"dog", "deer", "deal"},
		},
		out: []string{"deer", "deal"},
	},
	{
		desc: "substring",
		in: input{
			s: "de",
			a: []string{"dea", "deo", "dealo", "di"},
		},
		out: []string{"dea", "deo", "dealo"},
	},
	{
		desc: "empty",
		in: input{
			s: "foo",
			a: []string{},
		},
		out: []string{},
	},
	{
		desc: "no match",
		in: input{
			s: "foo",
			a: []string{"barfoo", "bar123"},
		},
		out: []string{},
	},
	{
		desc: "different starts",
		in: input{
			s: "foo",
			a: []string{"barfoo", "barfoo2", "latiza", "foo123"},
		},
		out: []string{"foo123"},
	},
	{
		desc: "empty pattern matches all",
		in: input{
			s: "",
			a: []string{"all", "strings", "go", "here"},
		},
		out: []string{"all", "strings", "go", "here"},
	},
	{
		desc: "repeated words in set",
		in: input{
			s: "123",
			a: []string{"123123", "123321", "123123"},
		},
		out: []string{"123123", "123321"},
	},
}

func testPrefixSearchFunc(prefixSearch func(s string, a []string) []string, t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			out := prefixSearch(tC.in.s, tC.in.a)

			sort.Strings(out)
			sort.Strings(tC.out)

			if !reflect.DeepEqual(tC.out, out) {
				t.Fatalf("expected %v, observed %v\n", tC.out, out)
			}
		})
	}
}

func TestPrefixSearch(t *testing.T) {
	testPrefixSearchFunc(PrefixSearch, t)
}

func TestPrefixSearch2(t *testing.T) {
	testPrefixSearchFunc(PrefixSearch2, t)
}
