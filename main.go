package main

import (
	"github.com/sshh12/apollgo/graphics"
	"github.com/sshh12/apollgo/web"
	"golang.org/x/mobile/app"
)

func main() {
	// go server.StartServer()
	go web.ServeSPA()
	app.Main(graphics.OnAppLaunch)
}
