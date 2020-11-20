package server

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/nadoo/glider/rule"
)

// ExternalIP is external IP
var ExternalIP = ""

// StartServer starts the glider server
func StartServer() {

	ExternalIP, _ = externalIP()

	rules := []*rule.Config{}
	stratCfg := &rule.StrategyConfig{}
	fowarders := []string{}

	pxy := rule.NewProxy(fowarders, stratCfg, rules)
	pxy.Check()

	local, err := NewMixedServer("tcp://:3080", pxy)
	if err != nil {
		log.Print(err)
		return
	}
	local.ListenAndServe()

}

func externalIP() (string, error) {
	resp, err := http.Get("https://ifconfig.me/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(body)), nil
}
