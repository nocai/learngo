package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filename := "config.yml"
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(file.Name())
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Printf(string(bytes))

}
