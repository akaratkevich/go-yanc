package actions

import "net"

// CidrToDecimalMask takes a *net.IPNet object and returns the decimal representation of the CIDR mask
func CidrToDecimalMask(ipNet *net.IPNet) string {
	mask := ipNet.Mask
	return net.IP(mask).String() // Convert mask to net.IP and then to string
}
