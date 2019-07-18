package t20190215

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func Test(t *testing.T) {
	b := NewBank()

	go b.Deposit("xiaoming", 100)
	go b.Withdraw("xiaoming", 20)
	go b.Deposit("xiaogang", 2000)

	time.Sleep(time.Second * 5)
	fmt.Printf("xiaoming has : %d\n", b.Query("xiaoming"))
	fmt.Printf("xiaogang has : %d\n", b.Query("xiaogang"))

}

type Bank struct {
	sync.Mutex
	saving map[string]int
}

func NewBank() *Bank {
	return &Bank{
		saving: make(map[string]int),
	}
}

func (b *Bank) Deposit(name string, amount int) {
	b.Lock()
	defer b.Unlock()

	if _, ok := b.saving[name]; !ok {
		b.saving[name] = 0
	}
	b.saving[name] += amount
}

func (b *Bank) Withdraw(name string, amount int) int {
	b.Lock()
	defer b.Unlock()

	if _, ok := b.saving[name]; !ok {
		return 0
	}

	if b.saving[name] < amount {
		amount = b.saving[name]
	}

	b.saving[name] -= amount

	return amount
}

func (b *Bank) Query(name string) int {
	b.Lock()
	defer b.Unlock()

	return b.saving[name]
}

