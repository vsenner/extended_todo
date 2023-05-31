package main

import (
	db "extended_todo/routing"
	"extended_todo/server"
)

func main() {
	db.SetupDB()
	server.Server()
}
