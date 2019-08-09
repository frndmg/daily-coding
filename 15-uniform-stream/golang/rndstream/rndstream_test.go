package rndstream

import (
	"math"
	"math/rand"
	"testing"
)

func sliceToChannel(s ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range s {
			out <- v
		}
		close(out)
	}()
	return out
}

func stats(N int, s []int, t *testing.T) (map[int]int, float64, float64) {
	data := make(map[int]int)

	for i := 0; i < N; i++ {
		ch := sliceToChannel(s...)
		data[Rnd(ch)]++
	}

	mean := func() float64 {
		sum := 0.0
		n := 0.0
		for _, v := range data {
			n++
			sum += float64(v)
		}
		return sum / n
	}

	variance := func(m float64) float64 {
		sum := 0.0
		n := 0.0
		for _, v := range data {
			n++
			sum += math.Pow(float64(v), 2)
		}
		return sum/n - math.Pow(m, 2)
	}

	m := mean()

	return data, m, math.Sqrt(variance(m))
}

func Test(t *testing.T) {
	// rand.Seed(int64(time.Now().Nanosecond()))
	rand.Seed(0)

	size := 50
	N := size * 1000
	s := make([]int, size)
	for i := 0; i < size; i++ {
		s[i] = i
	}

	_, mean, std := stats(N, s, t)

	// t.Log(data)
	t.Log("mean", mean, "variance", std)
}

var result int

func benchmark(size int, b *testing.B) {
	s := make([]int, size)

	var out int
	for i := 0; i < b.N; i++ {
		out = Rnd(sliceToChannel(s...))
	}
	result = out
}

func Benchmark1000(b *testing.B) {
	benchmark(1000, b)
}

func Benchmark10000(b *testing.B) {
	benchmark(10000, b)
}

func Benchmark100000(b *testing.B) {
	benchmark(100000, b)
}

func Benchmark1000000(b *testing.B) {
	benchmark(1000000, b)
}
