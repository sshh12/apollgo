package server

import (
	"fmt"

	"github.com/nadoo/glider/rule"
)

// StartServer starts the proxy server
func StartServer() {

	pxy := rule.NewProxy(
		[]string{},
		&rule.StrategyConfig{},
		[]*rule.Config{},
	)

	pxy.Check()

	local, err := NewMixedServer("tcp://:8443", pxy)
	if err != nil {
		fmt.Println(err)
		return
	}
	go local.ListenAndServe()

}
