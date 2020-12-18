package main

import (
	"fmt"
	gliderlog "github.com/nadoo/glider/log"
	"github.com/sshh12/apollgo/app"
	"github.com/sshh12/apollgo/network"
	"github.com/sshh12/apollgo/web"
)

func main() {
	apollgo := app.NewApollgoApp("apollgo.json")
	go apollgo.Run()
	apollgo.Log("Apollgo started.")
	go web.ServeWebApp(apollgo)
	gliderlog.F = func(s string, v ...interface{}) { apollgo.Log(fmt.Sprintf(s, v...)) }
	initCfg := apollgo.GetCfg()
	if initCfg.EnableGlider {
		apollgo.Log("Serving glider...")
		if err := network.ServeGlider(initCfg.Listeners); err != nil {
			apollgo.Log(err.Error())
		}
	} else {
		apollgo.Log("Glider disabled by settings.")
	}
	if initCfg.EnableHermes {
		apollgo.Log("Serving hermes...")
		if err := network.ServeHermes(initCfg.HermesConfig, apollgo.Log); err != nil {
			apollgo.Log(err.Error())
		}
	} else {
		apollgo.Log("Hermes disabled by settings.")
	}
	select {}
}
