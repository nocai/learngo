package main

import "fmt"

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 2)

	select {
	case <-ch1:
		fmt.Println("ch1")
	case <-ch2:
		fmt.Println("ch2")
	default:
		fmt.Println("default")
	}

	fmt.Println("over")
	//close(ch1)
	//close(ch2)
}
