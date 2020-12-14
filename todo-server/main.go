package main

import (
	"flag"
	"fmt"
	"os"

	"beards.ly/todo/controller"

	"beards.ly/todo/config"
	"beards.ly/todo/model"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)

	a := initializeApplication()
	a.DB.AutoMigrate(&model.User{})

	server := controller.NewServer(a)
	server.InitializeRoutes()
	server.Run(":8080")
}
