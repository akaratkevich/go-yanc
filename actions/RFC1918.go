package actions

import "net"

func RFC1918(ip net.IP) bool {
	privateNetworks := []string{"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16"}
	for _, network := range privateNetworks {
		_, ipNet, _ := net.ParseCIDR(network)
		if ipNet.Contains(ip) {
			return true
		}
	}
	return false
}
