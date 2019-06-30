package main

import (
	handlers "github.com/adrian-marcelo-gallardo/proxy-app/api/handlers"
	server "github.com/adrian-marcelo-gallardo/proxy-app/api/server"
	utils "github.com/adrian-marcelo-gallardo/proxy-app/api/utils"
)

func main() {
	utils.LoadEnv()
	app := server.SetUp()
	handlers.HandlerRedirection(app)
	server.RunServer(app)
}
