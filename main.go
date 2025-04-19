package main

import (
	"system/database"
	"system/server"
)

func main() {

	database.Connect()
	server.Handlers()
	server.ServerRun()

}