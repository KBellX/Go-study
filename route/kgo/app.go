package kgo

import (
	"fmt"
	"net/http"
)

type App struct {
	Handlers *ControllerRegister
	Server   *http.Server
}

func NewApp() *App {
	cr := NewControllerRegister()
	server := &http.Server{}
	return &App{Handlers: cr, Server: server}
}

func (app *App) Run() {
	app.Server.Addr = "127.0.0.1:8080"
	app.Server.Handler = app.Handlers

	endRunning := make(chan bool)

	go func() {
		fmt.Println("负责http.Server的协程跑起来了")
		err := app.Server.ListenAndServe()
		if err != nil {
			fmt.Println(err)
			endRunning <- true
		}
	}()
	<-endRunning

}
