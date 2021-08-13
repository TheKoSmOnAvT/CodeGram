package main

import (
	"backend/apiserver"
	"log"
)

func main() {

	config := apiserver.NewConfig()

	server := apiserver.New(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
