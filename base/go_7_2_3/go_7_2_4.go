package main

import "time"

func main() {
	w := make(chan bool)
	c := make(chan int, 2)
	
	go func() {
		select {
		case v := <- c:
			println(v)
		case <- time.After(time.Second):
			println("time out")
		}

		w <- true
	}()

	c <- 1
	println("wait")
	<- w
	println("done")
}
