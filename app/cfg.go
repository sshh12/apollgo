package app

// DefaultCfg default config
var DefaultCfg = &Config{
	Listeners: []ListenerConfig{
		ListenerConfig{
			URI:        "socks5://0.0.0.0:3080",
			Strategy:   "rr",
			Forwarders: []string{},
		},
	},
}

// Config server config
type Config struct {
	Listeners []ListenerConfig `json:"listeners"`
}

// ListenerConfig config for listeners
type ListenerConfig struct {
	URI        string   `json:"uri"`
	Strategy   string   `json:"strategy"`
	Forwarders []string `json:"forwarders"`
}
