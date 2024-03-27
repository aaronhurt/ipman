// Package local provides a machine local IP lookup backend
package local

import (
	"net"

	// backend definition
	"github.com/leprechau/ipman/common/ip"
)

// Get looks up client address
func (c *Config) Get(proto ip.IFlag) (string, error) {
	var conn net.Conn
	var err error

	// build connection based on requested proto
	switch proto {
	case ip.Inet:
		if conn, err = net.Dial("udp4", c.v4Host); err != nil {
			return "", err
		}
	case ip.Inet6:
		if conn, err = net.Dial("udp6", c.v6Host); err != nil {
			return "", err
		}
	}

	// close after return
	defer func() { _ = conn.Close() }()

	// get local address from connection object and return
	return conn.LocalAddr().(*net.UDPAddr).IP.String(), nil
}
