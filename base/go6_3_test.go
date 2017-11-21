package base

import (
	"fmt"
	"testing"
)

type User struct {
	id   int
	name string
}

func (self *User) String() string {
	return fmt.Sprintf("%d, %s", self.id, self.name)
}

func Test1(t *testing.T) {
	var o interface{} = &User{1, "Tome"}

	if i, ok := o.(fmt.Stringer); ok {
		fmt.Println(i)
	}

	u := o.(*User)
	//u := o.(User)
	fmt.Println(u)
}

func Test2(t *testing.T) {
	var o interface{} = &User{1, "Tom"}
	switch v := o.(type) {
	case nil:
		fmt.Println("nil")
	case fmt.Stringer:
		fmt.Println(v)
	case func() string:
		fmt.Println(v())
	case *User:
		fmt.Println("%d, %s", v.id, v.name)
	default:
		fmt.Println("unknown")
	}
}

