package server

import (
	"errors"
	"log"
	"net"

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
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("network error")
}
