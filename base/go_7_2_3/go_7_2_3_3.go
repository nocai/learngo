package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	var wg sync.WaitGroup
	quit := make(chan bool)

	for i := 0; i < 2; i ++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			task := func() {
				fmt.Println(id, time.Now().Nanosecond())
				time.Sleep(time.Second)
			}

			for {
				select {
				case <- quit:
					return
				default:
					task()
				}
			}
		}(i)
	}

	time.Sleep(time.Second * 6)

	close(quit)
	wg.Wait()
}
