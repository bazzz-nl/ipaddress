package ipaddress

import "net"

// Outbound returns the ip address used for outbound connections.
func Outbound() (net.IP, error) {
	conn, err := net.Dial("udp", "1.1.1.1:1") // Because it is udp, it does not actually connect.
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, nil
}
