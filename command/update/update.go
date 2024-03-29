// Package update provides the update cli command
package update

import (
	// backend definitions
	"github.com/leprechau/ipman/common/dns"
	"github.com/leprechau/ipman/common/ip"
)

// updateIP is a wrapper around the checkUpdate function
func (c *Command) updateIP() error {
	var err error

	// check local v4 address
	if c.config.v4 {
		if err = c.checkUpdate(ip.Inet, dns.A); err != nil {
			return err
		}
	}

	// check local v6 address
	if c.config.v6 {
		if err = c.checkUpdate(ip.Inet6, dns.AAAA); err != nil {
			return err
		}
	}

	// okay
	return nil
}

// checkUpdate checks the local machine IP and triggers a DNS update if needed
func (c *Command) checkUpdate(iType ip.IFlag, rType dns.RType) error {
	var local string
	var remote string
	var err error

	// get local address
	if local, err = c.ip.Get(iType); err != nil {
		return err
	}

	// get remote address
	if remote, err = c.dns.Get(c.config.domain, c.config.name, rType); err != nil {
		return err
	}

	// debugging
	c.Log.Printf("[%s] local/remote %s/%s", iType, local, remote)

	// update if needed
	if local != remote {
		// attempt to update remote record
		if err = c.dns.Upsert(c.config.domain, c.config.name, local, rType); err != nil {
			return err
		}

		// debugging
		c.Log.Printf("Updated remote %s record %s.%s -> %s",
			rType, c.config.name, c.config.domain, local)
	}

	// all okay
	return nil
}
