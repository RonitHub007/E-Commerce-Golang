package main

import (
	"E-Commerce-Golang/configs"
	"E-Commerce-Golang/internal/api"
	"fmt"
	"log"
)

func main() {
	fmt.Println("I am main function")
	config, err := configs.SetupEnv()

	if err != nil {
		log.Fatalln("config file is not there loaded properly")
	}
	api.StartServer(config)

}
