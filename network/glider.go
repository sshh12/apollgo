package network

import (
	"github.com/nadoo/glider/proxy"
	"github.com/nadoo/glider/rule"
	// proto support
	_ "github.com/nadoo/glider/proxy/http"
	_ "github.com/nadoo/glider/proxy/kcp"
	_ "github.com/nadoo/glider/proxy/mixed"
	_ "github.com/nadoo/glider/proxy/obfs"
	_ "github.com/nadoo/glider/proxy/reject"
	_ "github.com/nadoo/glider/proxy/socks4"
	_ "github.com/nadoo/glider/proxy/socks5"
	_ "github.com/nadoo/glider/proxy/ss"
	_ "github.com/nadoo/glider/proxy/ssh"
	_ "github.com/nadoo/glider/proxy/ssr"
	_ "github.com/nadoo/glider/proxy/tcp"
	_ "github.com/nadoo/glider/proxy/tls"
	_ "github.com/nadoo/glider/proxy/trojan"
	_ "github.com/nadoo/glider/proxy/udp"
	_ "github.com/nadoo/glider/proxy/vless"
	_ "github.com/nadoo/glider/proxy/vmess"
	_ "github.com/nadoo/glider/proxy/ws"
)

// ListenerConfig config for listeners
type ListenerConfig struct {
	URI           string   `json:"uri"`
	Strategy      string   `json:"strategy"`
	Forwarders    []string `json:"forwarders"`
	Check         string   `json:"check"`
	CheckInterval int      `json:"checkInterval"`
}

// ServeGlider starts glider server
func ServeGlider(listeners []ListenerConfig) error {
	for _, listener := range listeners {
		if err := runListener(&listener); err != nil {
			return err
		}
	}
	return nil
}

func runListener(listener *ListenerConfig) error {
	rules := []*rule.Config{}
	strat := &rule.Strategy{
		Strategy: listener.Strategy,
		Check:    listener.Check,
		CheckInterval: listener.CheckInterval,
	}
	forwarders := make([]string, 0)
	for _, forward := range listener.Forwarders {
		if forward != "" {
			forwarders = append(forwarders, forward)
		}
	}
	pxy := rule.NewProxy(forwarders, strat, rules)
	pxy.Check()
	local, err := proxy.ServerFromURL(listener.URI, pxy)
	if err != nil {
		return err
	}
	go local.ListenAndServe()
	return nil
}
