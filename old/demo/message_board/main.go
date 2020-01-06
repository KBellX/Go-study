package main

import (
	"./application"
)

func main() {
	server := &application.Manager{}
	server.Init()
}
