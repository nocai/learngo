package main

import (
	"github.com/spf13/viper"
	"log"
)

func main() {
	log.Println("aaaaa")
	log.Println("v")
	viper.AddConfigPath("a")
}
