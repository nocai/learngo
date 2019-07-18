package main

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:yjwggpl.@/test?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	//defer func() {
	//	if err := db.Close(); err != nil {
	//		panic(err)
	//	}
	//}()

	println("init the database")
	db.LogMode(true)
}

type Product struct {
	gorm.Model

	Code  string
	Price uint
}

func main() {
	db.AutoMigrate(&Product{})

	product := Product{Code: "code", Price: 1}
	//product.ID = 110
	db.Save(&product)
	log.Println("id = ", product.ID)
}
