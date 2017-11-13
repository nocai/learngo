package stack

import (
	"testing"
	"fmt"
)

func TestStack(t *testing.T) {
	var haystack Stack
	haystack.Push("hay")
	haystack.Push(-15)
	haystack.Push([]string{"pin", "clip", "needle"})
	haystack.Push(81.52)

	for {
		item, err := haystack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
