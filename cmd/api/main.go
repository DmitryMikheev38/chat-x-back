package main

import (
	"fmt"

	"github.com/dm/chat-x-back/config"
	"github.com/dm/chat-x-back/database"
	"github.com/dm/chat-x-back/server"
)

func main() {
	err := config.Init()
	if err != nil {
		panic(fmt.Sprintf("Error read config file: %+v", err.Error()))
	}

	err = database.Connect()
	if err != nil {
		panic(fmt.Sprintf("Error start postgres BD: %+v", err.Error()))
	}

	app := server.NewApp()
	app.Start()
}
