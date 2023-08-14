package actions

import (
	"fmt"
	"net"
)

func CidrToBinary(ip net.IP) string {
	binaryIP := ""
	for _, octet := range ip.To4() {
		binaryIP += fmt.Sprintf("%08b.", octet)
	}
	binaryIP = binaryIP[:len(binaryIP)-1] // Remove the trailing dot
	return binaryIP
}
