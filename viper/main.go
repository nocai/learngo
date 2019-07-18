package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("viper")
	viper.AddConfigPath("$HOME/.viper")
	viper.AddConfigPath("/Users/liujun/GolandProjects/learngo/viper")

	fmt.Println(viper.GetString("a.b"))
	fmt.Println("a")
}
