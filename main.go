package main

import (
	graphics "github.com/sshh12/apollgo/graphics"
	server "github.com/sshh12/apollgo/server"
	"golang.org/x/mobile/app"
)

func main() {
	go server.StartServer()
	app.Main(graphics.OnAppLaunch)
}
