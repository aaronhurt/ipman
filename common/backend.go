// Package common contains code shared between packages
package common

import (
	// ip backends
	"github.com/leprechau/ipman/common/ip"
	"github.com/leprechau/ipman/common/ip/ipify"
	"github.com/leprechau/ipman/common/ip/local"

	// dns backends
	"github.com/leprechau/ipman/common/dns"
	"github.com/leprechau/ipman/common/dns/godaddy"
)

// GetIPBackend returns an initialized IP lookup backend of the requested type
func GetIPBackend(backend string) (ip.Backend, error) {
	switch backend {
	case "ipify":
		return ipify.DefaultConfig()
	case "local":
		return local.DefaultConfig()
	}

	// if we got here it's a problem
	return nil, ErrUnknownIPBackend
}

// GetDNSBackend returns an initialized IP lookup backend of the requested type
func GetDNSBackend(backend string) (dns.Backend, error) {
	switch backend {
	case "godaddy":
		return godaddy.DefaultConfig()
	}

	// if we got here it's a problem
	return nil, ErrUnknownDNSBackend
}
