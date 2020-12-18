package network

import (
	"github.com/nadoo/glider/dns"
	"github.com/nadoo/glider/proxy"
	"github.com/nadoo/glider/rule"
	"net"

	"context"
	"time"
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
	URIs          []string `json:"uris"`
	Strategy      string   `json:"strategy"`
	Forwarders    []string `json:"forwarders"`
	Check         string   `json:"check"`
	CheckInterval int      `json:"checkInterval"`
	MaxFailures   int      `json:"maxFailures"`
	DialTimeout   int      `json:"dialTimeout"`
	RelayTimeout  int      `json:"relayTimeout"`
	IntFace       string   `json:"interface"`
	DNSListener   string   `json:"dns"`
	DNSAlwaysTCP  bool     `json:"dnsAlwaysTCP"`
	DNSServers    []string `json:"dnsServers"`
	DNSMaxTTL     int      `json:"dnsMaxTTL"`
	DNSMinTTL     int      `json:"dnsMinTTL"`
	DNSTimeout    int      `json:"dnsTimeout"`
	DNSCacheSize  int      `json:"dnsCacheSize"`
	DNSRecords    []string `json:"dnsRecords"`
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
		Strategy:      listener.Strategy,
		Check:         listener.Check,
		CheckInterval: listener.CheckInterval,
		MaxFailures:   listener.MaxFailures,
		DialTimeout:   listener.DialTimeout,
		RelayTimeout:  listener.RelayTimeout,
		IntFace:       listener.IntFace,
	}
	forwarders := make([]string, 0)
	for _, forward := range listener.Forwarders {
		if forward != "" {
			forwarders = append(forwarders, forward)
		}
	}
	pxy := rule.NewProxy(forwarders, strat, rules)
	if listener.DNSListener != "" {
		dnsConfig := &dns.Config{
			Servers:   listener.DNSServers,
			Timeout:   listener.DNSTimeout,
			MaxTTL:    listener.DNSMaxTTL,
			MinTTL:    listener.DNSMinTTL,
			Records:   listener.DNSRecords,
			AlwaysTCP: listener.DNSAlwaysTCP,
			CacheSize: listener.DNSCacheSize,
		}
		d, err := dns.NewServer(listener.DNSListener, pxy, dnsConfig)
		if err != nil {
			return err
		}
		d.AddHandler(pxy.AddDomainIP)
		d.Start()
		net.DefaultResolver = &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{Timeout: time.Second * 3}
				return d.DialContext(ctx, "udp", listener.DNSListener)
			},
		}
	}
	pxy.Check()
	for _, uri := range listener.URIs {
		if uri == "" {
			continue
		}
		local, err := proxy.ServerFromURL(uri, pxy)
		if err != nil {
			return err
		}
		go local.ListenAndServe()
	}
	return nil
}
