package logstore

import "testing"

var testCases = []struct {
	desc string
	cap  int
	ids  []OrderID
}{
	{
		desc: "minor cap",
		cap:  2,
		ids:  Range(5),
	},
	{
		desc: "cap - len = -1",
		cap:  4,
		ids:  Range(5),
	},
	{
		desc: "cap 1",
		cap:  1,
		ids:  Range(2),
	},
	{
		desc: "equal cap",
		cap:  2,
		ids:  Range(2),
	},
	{
		desc: "more cap",
		cap:  10,
		ids:  Range(5),
	},
}

func testLog(l Logger, cap int, d []OrderID, t *testing.T) {
	for _, id := range d {
		l.Record(id)
	}

	if len(d) < cap {
		cap = len(d)
	}

	for i := 1; i <= cap; i++ {
		id := l.GetLast(i)
		expected := d[len(d)-cap+i-1]
		if expected != id {
			t.Errorf("error log at index i=%v, expected=%v, observed=%v\n", i, expected, id)
		}
	}
}

func Test(t *testing.T) {
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			l := NewLog(tC.cap)
			testLog(l, tC.cap, tC.ids, t)
		})
	}
}

func Range(n int) []OrderID {
	out := make([]OrderID, n)
	for i := 0; i < n; i++ {
		out[i] = OrderID(i + 1)
	}
	return out
}
