package app

// DefaultCfg default config
var DefaultCfg = &Config{
	ApollgoPort: 8888,
	Listeners: []ListenerConfig{
		ListenerConfig{
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
	Listeners     []ListenerConfig `json:"listeners"`
	ApollgoPort   int              `json:"apollgoPort"`
	Check         string           `json:"check"`
	CheckInterval int              `json:"checkInterval"`
}

// ListenerConfig config for listeners
type ListenerConfig struct {
	URI        string   `json:"uri"`
	Strategy   string   `json:"strategy"`
	Forwarders []string `json:"forwarders"`
}
