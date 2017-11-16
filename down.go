package main

func vpnDown(cfg *Config) {
	printInfo("Disabling VPN killswitch rules.\n")

	executeUFW("--force", "reset")
	executeUFW("default", "deny", "incoming")
	executeUFW("default", "allow", "outgoing")
	executeUFW("enable")

	for _, app := range cfg.VPNapps {
		killApp(app)
	}
}
