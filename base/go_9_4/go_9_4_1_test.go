package go_9_4

import (
	"reflect"
	"testing"
	"fmt"
)

type User struct {
	Username string
}

type Admin struct {
	User
	title string
}

func Test(t *testing.T) {
	var u Admin
	tt := reflect.TypeOf(u)

	for i, n := 0, tt.NumField(); i < n; i ++ {
		f := tt.Field(i)
		fmt.Println(f.Name, f.Type)
	}
}

func Test2(a *testing.T) {
	u := new(Admin)

	t := reflect.TypeOf(u)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i, n := 0, t.NumField(); i < n; i ++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}
}
