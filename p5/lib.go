package p5

// Any type
type Any interface{}

// Pair type
type Pair func(func(Any, Any) Any) Any

// Cons creates a pair
func Cons(a, b Any) Pair {
	return func(f func(Any, Any) Any) Any {
		return f(a, b)
	}
}

// Car gets the first element of the pair
func Car(p Pair) Any {
	return p(func(a, _ Any) Any {
		return a
	})
}

// Cdr gets the second element of the pair
func Cdr(p Pair) Any {
	return p(func(_, b Any) Any {
		return b
	})
}
