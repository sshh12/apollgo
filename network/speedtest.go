package network

import (
	"time"

	"github.com/fopina/speedtest-cli/speedtest"
)

// RunSpeedTest runs a speedtest
func RunSpeedTest() (int, int, error) {
	client := speedtest.NewClient(&speedtest.Opts{
		Quiet:   true,
		Timeout: 30 * time.Second,
	})
	_, err := client.Config()
	if err != nil {
		return 0, 0, err
	}
	server, err := selectServer(client)
	if err != nil {
		return 0, 0, err
	}
	downloadSpeed := server.DownloadSpeed()
	uploadSpeed := server.UploadSpeed()
	return downloadSpeed * 8, uploadSpeed * 8, nil
}

func selectServer(client speedtest.Client) (*speedtest.Server, error) {
	servers, err := client.ClosestServers()
	if err != nil {
		return nil, err
	}
	selected := servers.MeasureLatencies(
		speedtest.DefaultLatencyMeasureTimes,
		speedtest.DefaultErrorLatency).First()
	return selected, nil
}
