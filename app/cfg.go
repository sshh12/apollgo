package app

import "github.com/sshh12/apollgo/network"

// DefaultCfg default config
var DefaultCfg = &Config{
	ApollgoPort: 8888,
	Listeners: []network.ListenerConfig{
		network.ListenerConfig{
			URIs:          []string{"socks5://0.0.0.0:3080"},
			Strategy:      "rr",
			Check:         "https://google.com",
			CheckInterval: 300,
			Forwarders:    []string{},
			MaxFailures:   3,
			DialTimeout:   3,
			RelayTimeout:  0,
			IntFace:       "",
			DNSListener:   "",
			DNSAlwaysTCP:  false,
			DNSServers:    []string{"8.8.8.8:53"},
			DNSMaxTTL:     1800,
			DNSMinTTL:     0,
			DNSTimeout:    3,
			DNSCacheSize:  4096,
			DNSRecords:    []string{},
		},
	},
}

// Config server config
type Config struct {
	Listeners   []network.ListenerConfig `json:"listeners"`
	ApollgoPort int                      `json:"apollgoPort"`
}
