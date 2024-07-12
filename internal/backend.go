// Package internal contains code shared between packages
package internal

import (
	// dns backends
	"github.com/leprechau/ipman/internal/dns"
	"github.com/leprechau/ipman/internal/dns/cloudflare"
	"github.com/leprechau/ipman/internal/dns/godaddy"
	"github.com/leprechau/ipman/internal/errors"
	"github.com/leprechau/ipman/internal/ip"
	"github.com/leprechau/ipman/internal/ip/ipify"
	"github.com/leprechau/ipman/internal/ip/local"
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
	return nil, errors.ErrUnknownIPBackend
}

// GetDNSBackend returns an initialized IP lookup backend of the requested type
func GetDNSBackend(backend string) (dns.Backend, error) {
	switch backend {
	case "godaddy":
		return godaddy.DefaultConfig()
	case "cloudflare":
		return cloudflare.DefaultConfig()
	}

	// if we got here it's a problem
	return nil, errors.ErrUnknownDNSBackend
}
