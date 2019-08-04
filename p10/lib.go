package p10

import (
	"time"
)

// After function
func After(n int, f func()) {
	after := time.AfterFunc(time.Duration(n)*time.Millisecond, f)
	go func() {
		<-after.C
	}()
}
