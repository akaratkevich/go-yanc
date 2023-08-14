package actions

import "net"

// CalculateHosts takes an IP network and returns the number of possible hosts
func CalculateHosts(ipNet *net.IPNet) int {
	ones, bits := ipNet.Mask.Size()
	return (1 << (bits - ones)) - 2 // Subtract 2 for network and broadcast addresses
}
