# killvpn - VPN kill switch toggler written in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/storvik/killvpn)](https://goreportcard.com/report/github.com/storvik/killvpn)

VPN kill switch that aims to prevent network traffic through anything other than specified interface.
Tested on Ubuntu 17.10 using Mullvad VPN.
[Inspiration and approach.](https://www.nukeador.com/06/07/2017/vpn-kill-switch-for-linux-protect-from-vpn-drops-and-dns-leaks/)

This VPN kill switch was not created by a security expert, nor can I guarnantee no leakage of anything.
Use at own risk!

## Usage

To run program, use `killvpn --config /path/to/config.json enable/disable`.

## Installation

Download, `make build`, move to desired location.

## Systemd

It's possible to have this run at boot etc. using systemd.
A systemd service file can be seen in `systemd/killvpn.service`.

## Configurations

See `killvpn.json` for configuration example.

- `VPNhosts` is an array with hostnames and ports.
- `VPNdevice` is the network device to be used with VPN. Check `ifconfig`.
- `NetworkDevices` is an array of network devices. Ex, eth0 for LAN and wlan0 for WLAN. Check `ifconfig`.
- `LocalNetworks` is an array of local area networks that should work. Can be an IP address or IP/subnet mask.
- `VPNapps` lists all apps that are NOT suppose to be running without VPN.
- `Verbose` true / false.
