package go_7_1

import (
	"testing"
	"sync"
	"runtime"
)

func Test2(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer println("A.defer")

		func() {
			defer println("B.defer")
			runtime.Goexit()
			println("B")
		}()
	}()

	wg.Wait()
}