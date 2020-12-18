package app

import (
	"encoding/json"
	"fmt"
	"github.com/sshh12/apollgo/network"
	"io/ioutil"
	"sync"
	"time"
)

// ApollgoApp is global app state
type ApollgoApp struct {
	cfg        *Config
	status     *Status
	cfgFn      string
	statusLock sync.Mutex
}

// LogLine is a single log
type LogLine struct {
	Text string `json:"text"`
	Time int64  `json:"time"`
}

// Status is app status
type Status struct {
	IP      string    `json:"ip"`
	DLSpeed float64   `json:"dlSpeed"`
	ULSpeed float64   `json:"ulSpeed"`
	Latency float64   `json:"latency"`
	Logs    []LogLine `json:"logs"`
}

// NewApollgoApp creates default state
func NewApollgoApp(cfgFn string) *ApollgoApp {
	cfg := DefaultCfg
	if cfgFile, err := ioutil.ReadFile(cfgFn); err == nil {
		var savedCfg Config
		if err := json.Unmarshal(cfgFile, &savedCfg); err == nil {
			cfg = &savedCfg
		} else {
			fmt.Println(err.Error())
		}
	}
	return &ApollgoApp{
		cfgFn: cfgFn,
		cfg:   cfg,
		status: &Status{
			IP:      "0.0.0.0",
			DLSpeed: 0,
			ULSpeed: 0,
			Latency: 0,
		},
		statusLock: sync.Mutex{},
	}
}

// Run starts misc tasks
func (s *ApollgoApp) Run() {
	ticker := time.NewTicker(2 * time.Hour)
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
func (s *ApollgoApp) Log(val string) {
	s.statusLock.Lock()
	defer s.statusLock.Unlock()
	newLog := LogLine{
		Time: time.Now().Unix(),
		Text: val,
	}
	s.status.Logs = append(s.status.Logs, newLog)
	if len(s.status.Logs) > 1000 {
		extra := 1000 - len(s.status.Logs)
		s.status.Logs = s.status.Logs[extra:]
	}
}

// GetCfg gets config
func (s *ApollgoApp) GetCfg() *Config {
	return s.cfg
}

// SetCfg updates config
func (s *ApollgoApp) SetCfg(newCfg *Config) {
	data, err := json.Marshal(newCfg)
	if err != nil {
		s.Log(err.Error())
		return
	}
	if err := ioutil.WriteFile(s.cfgFn, data, 0o644); err != nil {
		s.Log(err.Error())
		return
	}
	s.cfg = newCfg
}

// GetStatus gets status
func (s *ApollgoApp) GetStatus() *Status {
	s.statusLock.Lock()
	defer s.statusLock.Unlock()
	return s.status
}
