package main

import (
	handlers "github.com/edenriquez/proxy-app/api/handlers"
	server "github.com/edenriquez/proxy-app/api/server"
	utils "github.com/edenriquez/proxy-app/api/utils"
)

func main() {
	utils.LoadEnv()
	app := server.SetUp()
	handlers.HandlerRedirection(app)
	server.RunServer(app)
}
