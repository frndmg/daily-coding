package rndstream

import "math/rand"

// Rnd function
func Rnd(s <-chan int) int {
	b, ok := <-s
	if !ok {
		panic("at least one element")
	}

	for n := 2; ; n++ {
		if v, ok := <-s; ok {
			if rand.Intn(n) == 0 {
				b = v
			}
		} else {
			break
		}
	}

	return b
}
