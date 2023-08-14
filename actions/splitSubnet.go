package actions

import (
	"fmt"
	"net"
)

func SplitSubnet(network *net.IPNet, splitSize int) ([]string, error) {
	originalPrefixLength, _ := network.Mask.Size()
	if splitSize <= originalPrefixLength || splitSize > 32 {
		return nil, fmt.Errorf("invalid split size")
	}

	var subnets []string
	newMask := net.CIDRMask(splitSize, 32)
	ip := make(net.IP, len(network.IP))
	copy(ip, network.IP)

	for {
		subnet := &net.IPNet{IP: ip, Mask: newMask}
		subnets = append(subnets, subnet.String())

		// Increment IP by 2^(32 - splitSize)
		add := uint32(1) << (32 - splitSize)
		for i := len(ip) - 1; add > 0 && i >= 0; i-- {
			add += uint32(ip[i])
			ip[i] = byte(add & 0xff)
			add >>= 8
		}

		if !network.Contains(ip) {
			break
		}
	}

	return subnets, nil
}
