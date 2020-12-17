package main

import (
	"github.com/sshh12/apollgo/app"
	"github.com/sshh12/apollgo/web"
	mobileapp "golang.org/x/mobile/app"
)

func main() {
	apollgo := app.NewApollgoApp("/sdcard/apollgo.json")
	go apollgo.Run()
	go web.ServeWebApp(apollgo)
	// go server.StartServer()
	mobileapp.Main(app.OnAppLaunch)
}
