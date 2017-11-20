package go_7_2

import (
	"testing"
	"fmt"
	"os"
	"math/rand"
	"time"
	"sync"
)

func TestA(t *testing.T) {
	data := make(chan int)
	exit := make(chan bool)

	go func() {
		for d := range data {
			fmt.Println(d)
		}

		fmt.Println("recv over")
		exit <- true
	}()

	data <- 1
	fmt.Println("has send 1")
	data <- 2
	fmt.Println("has send 2")
	data <- 3
	fmt.Println("has send 3")
	close(data)

	fmt.Println("send over")
	<- exit
}


func TestB(t *testing.T) {
	data := make(chan int, 3)
	exit := make(chan bool)


	data <- 1
	data <- 2
	data <- 3
	go func() {
		for d := range data {
			fmt.Println(d)
		}
		exit <- true
	}()

	data <- 4
	data <- 5
	close(data)

	<- exit

}


func TestC(t *testing.T) {
	c := make(chan bool)
	cc := make(chan bool, 2)
	cc <- true
	fmt.Println(len(c), cap(c))
	fmt.Println(len(cc), cap(cc))

	var send chan<- bool = cc
	var recv <-chan bool = cc
	send <- true
	//<- send

	<- recv
	//recv <- true

	//d := (chan bool) (send)
}

func TestD(t *testing.T) {
	a, b := make(chan int, 3), make(chan int)

	go func() {
		v, ok, s := 0, false, ""

		for {
			select {
			case v, ok = <-a:
				s = "a"
			case v, ok = <-b:
				s = "b"
			}

			if ok {
				fmt.Println(v, s)
			} else {
				os.Exit(0)
			}

		}
	}()

	for i := 0; i < 5; i ++ {
		select {
		case a<-i:
		case b <-i:
		}
	}

	close(a)
	select {}
}

func NewTest() chan int {
	c := make(chan int)
	rand.Seed(time.Now().UnixNano())

	go func() {
		time.Sleep(time.Second)
		c <- rand.Int()
	}()
	return c
}

func TestE(t *testing.T) {
	newTest := NewTest()
	println(<-newTest)
}

func TestF(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(3)

	sem := make(chan int, 1)

	for i:= 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()

			//sem := make(chan int, 1)
			sem <- 1

			for x := 0; x < 3; x ++ {
				fmt.Println(id, x)
			}
			<- sem
		}(i)
	}
	wg.Wait()
}
