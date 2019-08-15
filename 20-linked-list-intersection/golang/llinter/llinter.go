package llinter

// LList type
type LList struct {
	val  int
	next *LList
}

// Intersection function
func Intersection(a, b *LList) *LList {
	m := Len(a)
	n := Len(b)

	if n > m {
		m, n = n, m
		a, b = b, a
	}

	// Remove m - n elements from a
	for i := 0; i < (m - n); i++ {
		a = a.next
	}

	for a != nil && b != nil {
		if a.val == b.val {
			return a
		}
		a = a.next
		b = b.next
	}

	return nil
}

// Len function
func Len(a *LList) int {
	n := 0
	for a != nil {
		n++
		a = a.next
	}
	return n
}
