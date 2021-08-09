package main

import (
	"backend/database"
)

func main() {
	database.InitDB()
	database.Login()
}
