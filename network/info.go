package network

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// ExternalIP gets external ip
func ExternalIP() (string, error) {
	resp, err := http.Get("https://ifconfig.me/")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(body)), nil
}
