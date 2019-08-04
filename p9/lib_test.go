package p9

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc string
		in   []int
		out  int
	}{
		{
			desc: "example 1",
			in:   []int{2, 4, 6, 2, 5},
			out:  13,
		},
		{
			desc: "example 2",
			in:   []int{5, 1, 1, 5},
			out:  10,
		},
		{
			desc: "two elements",
			in:   []int{4, 5},
			out:  5,
		},
		{
			desc: "two elements II",
			in:   []int{5, 4},
			out:  5,
		},
		{
			desc: "",
			in:   []int{2, 2, 3, 4, 5, 6},
			out:  12,
		},
		{
			desc: "",
			in:   []int{6, 5, 4, 3, 2, 2},
			out:  12,
		},
		{
			desc: "one element",
			in:   []int{1},
			out:  1,
		},
		{
			desc: "no elements",
			in:   []int{},
			out:  0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			out := MaxSum(tC.in)
			if !reflect.DeepEqual(tC.out, out) {
				t.Fatalf("expected %v, observed %v\n", tC.out, out)
			}
		})
	}
}

// random returns a random in [a, b] range
func random(r *rand.Rand, a, b int) int {
	return r.Intn(b-a) + a
}

func TestRandom(t *testing.T) {
	N := 5000

	L := 1000
	m, M := -1000, 1000

	// s := rand.NewSource(int64(time.Now().Nanosecond()))
	s := rand.NewSource(42)
	r := rand.New(s)

	for i := 0; i < N; i++ {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			L := random(r, 0, L)

			in := make([]int, L)
			for i := 0; i < L; i++ {
				in[i] = random(r, m, M)
			}

			inC := make([]int, L)

			copy(inC, in)
			out := MaxSum(inC)

			copy(inC, in)
			outSlow := MaxSumSlow(inC)

			if !reflect.DeepEqual(out, outSlow) {
				t.Fatalf("in: %v, out: %v, outSlow: %v\n", in, out, outSlow)
			}
		})
	}
}

var result int

func benchmarkMaxSumFunc(L, m, M int, r *rand.Rand, maxSum func([]int) int, b *testing.B) {
	in := make([]int, L)
	for j := 0; j < L; j++ {
		in[j] = random(r, m, M)
	}

	var out int
	for i := 0; i < b.N; i++ {
		out = maxSum(in)
	}
	result = out
}

func benchmarkMaxSumFuncParallel(L, m, M int, r *rand.Rand, maxSum func([]int) int, b *testing.B) {
	in := make([]int, L)
	for j := 0; j < L; j++ {
		in[j] = random(r, m, M)
	}

	var out int
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			out = maxSum(in)
		}
	})
	result = out
}

func benchmarkMaxSum(L, m, M int, r *rand.Rand, b *testing.B) {
	benchmarkMaxSumFunc(L, m, M, r, MaxSum, b)
}

func benchmarkMaxSumParallel(L, m, M int, r *rand.Rand, b *testing.B) {
	benchmarkMaxSumFuncParallel(L, m, M, r, MaxSum, b)
}

func benchmarkMaxSumSlow(L, m, M int, r *rand.Rand, b *testing.B) {
	benchmarkMaxSumFunc(L, m, M, r, MaxSumSlow, b)
}

func newRandom() *rand.Rand {
	return rand.New(rand.NewSource(42))
}

func BenchmarkMaxSum1000(b *testing.B) {
	benchmarkMaxSum(1000, -1000, 1000, newRandom(), b)
}

func BenchmarkMaxSum10000(b *testing.B) {
	benchmarkMaxSum(10000, -1000, 1000, newRandom(), b)
}

func BenchmarkMaxSum100000(b *testing.B) {
	benchmarkMaxSum(100000, -1000, 1000, newRandom(), b)
}

func BenchmarkMaxSum1000000(b *testing.B) {
	benchmarkMaxSum(1000000, -1000, 1000, newRandom(), b)
}

func BenchmarkMaxSumSlow1000(b *testing.B) {
	benchmarkMaxSumSlow(1000, -1000, 1000, newRandom(), b)
}

func BenchmarkMaxSumSlow5000(b *testing.B) {
	benchmarkMaxSumSlow(5000, -1000, 1000, newRandom(), b)
}

func BenchmarkMaxSumSlow10000(b *testing.B) {
	benchmarkMaxSumSlow(10000, -1000, 1000, newRandom(), b)
}
