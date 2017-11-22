package go_7_4_2

import (
	"testing"
	"reflect"
	"fmt"
)

type User struct {

}

type Admin struct {
	User
}

func (*User) ToString() {}

func (Admin) test() {}

func Test(tt *testing.T) {
	var u Admin

	methods := func(t reflect.Type) {
		for i, n := 0, t.NumMethod(); i < n; i ++ {
			m := t.Method(i)
			fmt.Println(m.Name)
		}
	}

	fmt.Println("--------- value interface ------------")
	methods(reflect.TypeOf(u))

	fmt.Println("---------------prointer interface --------------")
	methods(reflect.TypeOf(&u))
}