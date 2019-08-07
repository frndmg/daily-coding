package p11

import "strings"

// PrefixSearch function
func PrefixSearch(s string, possible []string) []string {
	out := make([]string, 0)
	for _, p := range possible {
		if strings.HasPrefix(p, s) {
			out = append(out, p)
		}
	}
	return out
}
