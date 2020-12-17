package main

import (
	"github.com/sshh12/apollgo/app"
	"github.com/sshh12/apollgo/web"
	mobileapp "golang.org/x/mobile/app"
)

func main() {
	state := app.NewAppState()
	go state.Run()
	go web.ServeWebApp(state)
	// go server.StartServer()
	mobileapp.Main(app.OnAppLaunch)
}
