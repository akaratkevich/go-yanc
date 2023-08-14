package main

import (
	"flag"
	"fmt"
	"github.com/pterm/pterm"
	"go/yanc/actions"
	"net"
	"os"
	"strings"
)

func main() {
	// Overwrite DefaultHeader style
	pterm.DefaultHeader = *pterm.DefaultHeader.WithBackgroundStyle(pterm.NewStyle(pterm.BgDefault))
	pterm.DefaultHeader = *pterm.DefaultHeader.WithMargin(30)
	pterm.DefaultHeader = *pterm.DefaultHeader.WithTextStyle(pterm.NewStyle(pterm.FgLightYellow, pterm.Bold))
	pterm.DefaultHeader.Println("* IPV4 Network Details: *")

	// Command line flags
	// Define a flag for the network string
	networkStr := flag.String("network", "", "IPv4 network in CIDR format (e.g., 192.168.1.0/24)")
	split := flag.String("split", "", "Split network into subnets of this size (e.g., /25)")

	flag.Parse() // Parse the command-line flags

	// Check that the network string has been provided
	if *networkStr == "" {
		fmt.Println("Please provide the network using the --network flag")
		os.Exit(1)
	}

	// Validate the network string
	ip, ipNet, err := net.ParseCIDR(*networkStr)
	if err != nil {
		fmt.Println("Invalid network:", err)
		os.Exit(1)
	}
	if actions.RFC1918(ip) {
		pterm.Info.Println("The provided subnet is part of the RFC1918 private address space.\n")
	} else {
		pterm.Info.Println("The provided subnet is NOT part of the RFC1918 private address space.\n")
	}

	networkPrts := strings.Split(*networkStr, "/")
	if len(networkPrts) != 2 {
		fmt.Println("Invalid network string")
		return
	}

	binaryIP := actions.CidrToBinary(ip)
	decimalMask := actions.CidrToDecimalMask(ipNet)
	binaryMask := actions.CidrMaskToBinary(networkPrts[1])
	numberOfHosts := actions.CalculateHosts(ipNet)

	// Create panel 1
	//panel1 := pterm.FgGray.Sprintf("\nCIDR: %s\nDecimal Mask: %s\nBinary IP: %s\nBinary Mask: %s\nNumber of Hosts: %d\n", *networkStr, decimalMask, binaryIP, binaryMask, numberOfHosts)
	panel1 := pterm.FgGray.Sprintf("├─ CIDR: %s", *networkStr)
	panel1 += pterm.FgGray.Sprintf("\n│   ├─ Decimal Mask: %s", decimalMask)
	panel1 += pterm.FgGray.Sprintf("\n│   ├─ Binary IP: %s", binaryIP)
	panel1 += pterm.FgGray.Sprintf("\n│   ├─ Binary Mask: %s", binaryMask)
	panel1 += pterm.FgGray.Sprintf("\n│   └─ Number of Hosts: %d", numberOfHosts)

	// Panel2 is for the subnet split option
	panel2 := ""
	// Parse the split size
	if *split != "" {
		splitSize := 0
		if _, err := fmt.Sscanf(*split, "/%d", &splitSize); err != nil {
			fmt.Println("Invalid split size:", err)
			os.Exit(1)
		}

		// Split the subnets
		splitSubnets, err := actions.SplitSubnet(ipNet, splitSize)
		if err != nil {
			fmt.Println("Error splitting subnets:", err)
			os.Exit(1)
		}
		// Calculate number of subnets for the split
		numberOfNetworks := len(splitSubnets)

		// Calculate number of hosts per network by parsing the first subnet
		_, firstSubnetNet, err := net.ParseCIDR(splitSubnets[0])
		if err != nil {
			fmt.Println("Error parsing subnet:", err)
			os.Exit(1)
		}
		numberOfHostsPerNetwork := actions.CalculateHosts(firstSubnetNet)
		firstSubnetDecimalMask := actions.CidrToDecimalMask(firstSubnetNet)

		// Create panel 2
		panel2 += pterm.FgLightGreen.Sprintf("Split into %d networks, %d hosts per network\n", numberOfNetworks, numberOfHostsPerNetwork)
		panel2 += pterm.FgLightGreen.Sprintf("Decimal Mask: %s\n", firstSubnetDecimalMask)
		panel2 += pterm.FgGray.Sprintf("\n├─ %s", *networkStr) // Root of the tree
		for i, subnet := range splitSubnets {
			if i < len(splitSubnets)-1 {
				panel2 += pterm.FgGray.Sprintf("\n│   ├─ " + subnet) // Branches of the tree
			} else {
				panel2 += pterm.FgGray.Sprintf("\n│   └─ " + subnet) // Last branch of the tree
			}
		}
	}
	// Render  panels using DefaultBox
	panels := pterm.Panels{
		{{Data: pterm.DefaultBox.WithTitle("Current Network").Sprintf(panel1)}, {Data: pterm.DefaultBox.WithTitle("Requested Split").Sprintf(panel2)}},
	}
	// Print panels.
	pterm.DefaultPanel.WithPanels(panels).Render()
	// End panels
}
