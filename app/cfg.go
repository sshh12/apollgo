package app

import "github.com/sshh12/apollgo/network"

// DefaultCfg default config
var DefaultCfg = &Config{
	ApollgoPort: 8888,
	Listeners: []network.ListenerConfig{
		network.ListenerConfig{
			URI:           "socks5://0.0.0.0:3080",
			Strategy:      "rr",
			Check:         "https://google.com",
			CheckInterval: 300,
			Forwarders:    []string{},
		},
	},
}

// Config server config
type Config struct {
	Listeners   []network.ListenerConfig `json:"listeners"`
	ApollgoPort int                      `json:"apollgoPort"`
}
