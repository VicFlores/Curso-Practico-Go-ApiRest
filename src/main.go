package main

import (
	Handle "github.com/VicFlores/src/handlers"
	"github.com/VicFlores/src/middlewares"
	Server "github.com/VicFlores/src/server"
)

func main() {

	server := Server.NewServer(":4000")
	server.Handle("/", "GET", server.AddMiddleware(Handle.HandleHome, middlewares.CheckAuth(), middlewares.Logging()))
	server.Handle("/shop", "POST", server.AddMiddleware(Handle.HandleShop, middlewares.CheckAuth(), middlewares.Logging()))
	server.Handle("/create-user", "POST", Handle.PostUser)
	server.Listen()

}
