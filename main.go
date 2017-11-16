package main

import (
	"flag"
)

// Config struct containing firewall rules parameters
type Config struct {
	VPNhosts []struct {
		Hostname string `json:"Hostname"`
		Port     string `json:"Port"`
	} `json:"VPNhosts"`
	VPNdevice      string   `json:"VPNdevice"`
	NetworkDevices []string `json:"NetworkDevices"`
	LocalNetworks  []string `json:"LocalNetworks"`
	VPNapps        []string `json:"VPNapps"`
}

var verbose bool

func main() {
	// Parse flags
	cfgPath := flag.String("config", "~/.killvpn.json", "path to configuration file")
	vrbse := flag.Bool("verbose", false, "verbose output")
	flag.Parse()
	verbose = *vrbse

	// Parse config file
	config, err := readConfig(*cfgPath)
	if err != nil {
		panic(err)
	}

	if len(flag.Args()) == 0 {
		vpnUp(config)
		return
	}

	if flag.Args()[0] == "disable" {
		vpnDown(config)
		return
	}

	vpnUp(config)
}
