package main

import (
	"backend/apiserver"
	"log"
)

var (
//	configPath string
)

func init() {
	//prblems with read .toml files
	//flag.StringVar(&configPath, "config-path", "config/apiserver.toml", "path to config file")
}

func main() {
	//flag.Parse()

	config := apiserver.NewConfig()
	// _, err := toml.DecodeFile(configPath, config)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	server := apiserver.New(config)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	//database.InitDB()
	//database.Login()
}
