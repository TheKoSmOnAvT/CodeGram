package database

import (
	"backend/dbModels"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db gorm.DB

func InitDB() {

	const pathToDB string = "../db/Codegram.db"

	_, err := gorm.Open(sqlite.Open(pathToDB), &gorm.Config{})
	if err != nil {
		print(err)
		panic("failed to connect database")
	}
	fmt.Print("OK")
}

func Login() {
	var acc dbModels.Account
	res := db.First(&acc, 1) // find product with integer primary key
	fmt.Printf("%+v", res)
}
