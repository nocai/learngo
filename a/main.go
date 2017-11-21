package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string = "abcdefghijklmn我们也是一样"
	fmt.Println(str[9:])
	fmt.Println(len(str))
	fmt.Println([]rune(str))
	fmt.Println(string([]rune(str)))

	fmt.Println(len([]rune(str)))

	fmt.Println([]byte(str))
	fmt.Println(len([]byte(str)))

	fmt.Println(string(str[1]))

	for _, s := range str {
		fmt.Println(reflect.TypeOf(s))
		fmt.Println(s)
	}
}
