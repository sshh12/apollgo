package main

import (
	"fmt"
	gliderlog "github.com/nadoo/glider/log"
	"github.com/sshh12/apollgo/app"
	"github.com/sshh12/apollgo/network"
	"github.com/sshh12/apollgo/web"
	mobileapp "golang.org/x/mobile/app"
)

func main() {
	apollgo := app.NewApollgoApp("/sdcard/apollgo.json")
	go apollgo.Run()
	go web.ServeWebApp(apollgo)
	gliderlog.F = func(s string, v ...interface{}) { apollgo.Log(fmt.Sprintf(s, v...)) }
	initCfg := apollgo.GetCfg()
	if err := network.ServeGlider(initCfg.Listeners); err != nil {
		apollgo.Log(err.Error())
	}
	mobileapp.Main(app.OnAppLaunch)
}
