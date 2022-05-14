package maths

import (
	"fmt"
	"testing"
)

func TestRandom(t *testing.T) {
	for i := 0; i < 10000; i++ {
		fmt.Printf("%d,", Random(2, 4))
	}
}
