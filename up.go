package main

import (
	"net"
)

func vpnUp(cfg *Config) {
	printInfo("Enabling VPN killswitch rules.\n")

	// Reset and set default rules
	executeUFW("--force", "reset")
	executeUFW("default", "deny", "incoming")
	executeUFW("default", "deny", "outgoing")

	portMap := map[string]string{}

	// Allow communication to all VPN hosts
	for _, vpnHost := range cfg.VPNhosts {
		tmp := net.ParseIP(vpnHost.Hostname)
		if tmp.To4() == nil {
			// Hostname not valid IP, need to use lookup
			hosts, err := net.LookupHost(vpnHost.Hostname)
			if err != nil {
				panic(err)
			}
			for _, host := range hosts {
				for _, device := range cfg.NetworkDevices {
					executeUFW("allow", "out", "on", device, "from", "any", "to", host)
				}
			}
		} else {
			// Hostname already valid IP
			for _, device := range cfg.NetworkDevices {
				executeUFW("allow", "out", "on", device, "from", "any", "to", vpnHost.Hostname)
			}
		}
		// Add port to portMap
		if _, ok := portMap[vpnHost.Port]; !ok {
			printInfo("Adding port %s to portMap\n", vpnHost.Port)
			portMap[vpnHost.Port] = vpnHost.Port
		}
	}

	// Allow out on all vpn ports and DNS port
	for _, port := range portMap {
		executeUFW("allow", "out", port+"/udp")
	}
	executeUFW("allow", "out", "53/udp")

	// Allow all communication through VPN tunnel device
	executeUFW("allow", "out", "on", cfg.VPNdevice, "from", "any", "to", "any")
	executeUFW("allow", "in", "on", cfg.VPNdevice, "from", "any", "to", "any")

	// Allow communication to LAN
	for _, device := range cfg.NetworkDevices {
		for _, lan := range cfg.LocalNetworks {
			executeUFW("allow", "out", "on", device, "from", "any", "to", lan)
			executeUFW("allow", "in", "on", device, "from", lan, "to", "any")
		}
	}

	// Enable firewall
	executeUFW("enable")

}
