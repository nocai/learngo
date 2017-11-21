package base

import (
	"fmt"
	"testing"
)

type Stringer interface {
	String() string
}

type Printer interface {
	String() string
	Print()
}


func (this *User) Print() {
	fmt.Println(this.String())
}

func Test(t *testing.T) {
	var o Printer = &User{1, "tome"}
	var s Stringer = o
	fmt.Println(s.String())
	o.Print()

	var _ fmt.Stringer = (*User)(nil)
}

type Tester interface {
	Do()
}

type FuncDo func()
func (this FuncDo) Do() {
	this()
}

func Test11(a *testing.T) {
	var fc = func() {
		fmt.Println("Hello World")
	}
	var t Tester = FuncDo(fc)
	t.Do()
}