package go_7_1

import (
	"math"
	"sync"
	"testing"
)

func sum(id int) {
	var x int64
	for i := 0; i<math.MaxUint32; i++ {
		x += int64(i)
	}
	print(id, x)
}

func Test(t *testing.T) {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	for i := 0; i < 2; i++ {
		go func(id int) {
			defer wg.Done()
			sum(id)
		}(i)
	}

	wg.Wait()
}