package p2

import "errors"

var (
	errSmallList = errors.New("p2: unable to handle array of size less or equal than 1")
)

func prefix(l []int) []int {
	p := make([]int, len(l))
	p[0] = l[0]
	for i := 1; i < len(l); i++ {
		p[i] = p[i-1] * l[i]
	}
	return p
}

func sufix(l []int) []int {
	s := make([]int, len(l))
	s[len(l)-1] = l[len(l)-1]
	for i := len(l) - 2; i >= 0; i-- {
		s[i] = s[i+1] * l[i]
	}
	return s
}

// MultExcepti function
func MultExcepti(l []int) ([]int, error) {
	n := len(l)

	if n <= 1 {
		return nil, errSmallList
	}

	p := prefix(l)
	s := sufix(l)

	res := make([]int, n)
	res[0] = s[1]
	res[n-1] = p[n-2]
	for i := 1; i < n-1; i++ {
		res[i] = p[i-1] * s[i+1]
	}

	return res, nil
}
