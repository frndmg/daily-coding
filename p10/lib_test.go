package p10

import (
	"fmt"
	"testing"
	"time"
)

func Test(t *testing.T) {
	After(1000, func() {
		fmt.Println("Hello After")
	})

	time.Sleep(2 * time.Second)
}
