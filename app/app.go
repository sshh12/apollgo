package app

import (
	"github.com/sshh12/apollgo/network"
	"log"
	"time"
)

// State is global app state
type State struct {
	cfg    *Config
	status *Status
}

// Status is app status
type Status struct {
	IP      string  `json:"ip"`
	DLSpeed float64 `json:"dlSpeed"`
	ULSpeed float64 `json:"ulSpeed"`
	Latency float64 `json:"latency"`
}

// NewAppState creates default state
func NewAppState() *State {
	return &State{
		cfg: DefaultCfg,
		status: &Status{
			IP:      "0.0.0.0",
			DLSpeed: 0,
			ULSpeed: 0,
			Latency: 0,
		},
	}
}

// Run starts misc tasks
func (s *State) Run() {
	ticker := time.NewTicker(300 * time.Second)
	for ; true; <-ticker.C {
		if ip, err := network.ExternalIP(); err == nil {
			s.status.IP = ip
		} else {
			s.Log(err.Error())
		}
		if dl, up, lat, err := network.RunSpeedTest(); err == nil {
			s.status.DLSpeed = dl
			s.status.ULSpeed = up
			s.status.Latency = lat
		} else {
			s.Log(err.Error())
		}
	}
}

// Log logs something
func (s *State) Log(val string) {
	log.Println("[APOLLGO] " + val)
}

// GetCfg gets config
func (s *State) GetCfg() *Config {
	return s.cfg
}

// GetStatus gets status
func (s *State) GetStatus() *Status {
	return s.status
}
