package network

import (
	"fmt"
	"net"
	"strings"
	"strconv"
	hermesTCP "github.com/sshh12/hermes/tcp"
)

// HermesConfig configures hermes settings
type HermesConfig struct {
	Password     string   `json:"password"`
	HermesPort   int   `json:"port"`
	Server       string   `json:"server"`
	ForwardPairs []string `json:"forwardPairs"`
}

// ServeHermes starts hermes client
func ServeHermes(cfg HermesConfig, log func(string)) error {
	for _, pair := range cfg.ForwardPairs {
		split := strings.Split(pair, "/")
		if len(split) != 2 {
			return fmt.Errorf("Invalid port pair %s", pair)
		}
		appPort, err := strconv.Atoi(split[0])
		if err != nil {
			return err
		}
		remotePort, err := strconv.Atoi(split[1])
		if err != nil {
			return err
		}
		serverAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", cfg.Server, cfg.HermesPort))
		if err != nil {
			return err
		}
		go func() {
			client, err := hermesTCP.NewClient(
				appPort, 
				remotePort, 
				cfg.Server, 
				hermesTCP.WithRestarts(), 
				hermesTCP.WithPassword(cfg.Password),
				hermesTCP.WithServerAddress(serverAddr),
			)
			if err != nil {
				log(err.Error())
			}
			if err := client.Start(); err != nil {
				log(err.Error())
			}
		}()
	}
	return nil
}