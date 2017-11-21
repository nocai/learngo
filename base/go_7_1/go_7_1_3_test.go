package go_7_1

import (
	"testing"
	"sync"
	"runtime"
)

func Test3(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		for i := 0; i < 5; i++ {
			println(i)
			if i == 1 {
				runtime.Gosched()
			}
		}
	}()

	go func() {
		defer wg.Done()
		println("Hello world")
	}()

	wg.Wait()
}
