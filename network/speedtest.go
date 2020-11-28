package network

import (
	"time"
	"fmt"
	"sort"
	"github.com/showwin/speedtest-go/speedtest"
)

// RunSpeedTest runs a speedtest
func RunSpeedTest() (float64, float64, float64, error) {
	user, err:= speedtest.FetchUserInfo()
	if err != nil {
		return 0, 0, 0, err
	}
	serverList, err := speedtest.FetchServerList(user)
	if err != nil {
		return 0, 0, 0, err
	}
	targets,  err := serverList.FindServer([]int{})
	if err != nil {
		return 0, 0, 0, err
	}
	if len(targets) == 0 {
		return 0, 0, 0, fmt.Errorf("no speed test servers found")
	}
	sort.Slice(targets, func(i, j int) bool {
		return targets[i].Distance > targets[j].Distance
	})
	s := targets[0]
	s.PingTest()
	s.DownloadTest()
	s.UploadTest()
	return s.DLSpeed, s.ULSpeed, float64(s.Latency / time.Millisecond), nil
}