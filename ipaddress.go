package ipaddress

import (
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
)

// Outbound returns the ip address used for outbound connections.
func Outbound() (string, error) {
	conn, err := net.Dial("udp", "1.1.1.1:1") // Because it is udp, it does not actually connect.
	if err != nil {
		return "", err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

// Public returns the public ip address for outbound connections behind NAT.
func Public() (string, error) {
	response, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	type ip struct {
		Status string `json:"status"`
		Query  string `json:"query"`
	}
	address := ip{}
	err = json.Unmarshal(body, &address)
	if err != nil {
		return "", err
	}
	if address.Status != "success" {
		return "", errors.New("could not get ip adress")
	}
	return address.Query, nil
}
