package justtext

import "strings"

// Justify function
func Justify(ws []string, k int) []string {
	out := []string{}

	line := []string{}
	lineCount := 0

	for i := 0; i < len(ws); i++ {
		w := ws[i]

		if lineCount+len(w)+len(line) > k {
			out = append(out, makeFlatLine(line, k))
			line = []string{w}
			lineCount = len(w)
			continue
		}

		line = append(line, w)
		lineCount += len(w)

		if i == len(ws)-1 {
			out = append(out, makeFlatLine(line, k))
			break // it should happen anyway
		}
	}

	return out
}

func makeFlatLine(line []string, k int) string {
	lineCount := 0
	for _, w := range line {
		lineCount += len(w)
	}

	if len(line) == 0 {
		return strings.Repeat(" ", k)
	}
	if len(line) == 1 {
		return line[0] + strings.Repeat(" ", k-len(line[0]))
	}

	n := len(line) - 1
	t := k - lineCount

	m := t / n
	r := t % n

	// m * n + r = t

	//    (m+1)*r + m*(n-r) = t
	// => m*r + r + m*n - m*r = t
	// => m*n + r = t
	// :ok:

	left := strings.Join(line[:r+1], strings.Repeat(" ", m+1))
	right := strings.Join(line[r:], strings.Repeat(" ", m))

	return left + right[len(line[r]):]
}
