package actions

import "net"

func CalculateFirstLastBroadcast(ipNet *net.IPNet) (net.IP, net.IP, net.IP, net.IP) {
	mask := ipNet.Mask
	network := ipNet.IP.Mask(mask)
	firstUsable := net.IP(make([]byte, len(network)))
	lastAddr := net.IP(make([]byte, len(network)))
	broadcast := net.IP(make([]byte, len(network)))

	for i := 0; i < len(network); i++ {
		firstUsable[i] = network[i]
		lastAddr[i] = network[i] | ^mask[i]
		broadcast[i] = network[i] | ^mask[i]
	}

	// Handle /32 subnet
	ones, bits := mask.Size()
	if ones == bits { // If all bits are set, it's a /32 subnet
		firstUsable = network
		lastAddr = network
		broadcast = network
	}

	// Increment firstUsable and decrement lastAddr to get last usable address
	firstUsable[len(firstUsable)-1]++
	lastAddr[len(lastAddr)-1]--

	return network, firstUsable, lastAddr, broadcast
}
