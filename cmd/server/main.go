package main

import (
	"fmt"
	"pet-paradise/api"
	"pet-paradise/config"
	"pet-paradise/log"
)

func main() {
	fmt.Println(VERSION)
	fmt.Println("Pet-Paradise Server")
	if err := config.ParseConfig("./config.yaml"); err != nil {
		fmt.Printf("%#v", err)
		return
	}
	log.Logger().Info("start server")

	r := api.InitRouter()
	r.Run(":8080")
}
