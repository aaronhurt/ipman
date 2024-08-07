// Package update provides the update cli command
package update

import (
	// backend definitions
	"github.com/leprechau/ipman/internal/dns"
	"github.com/leprechau/ipman/internal/ip"
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

	return nil
}

// checkUpdate checks the local machine IP and triggers a DNS update if needed
func (c *Command) checkUpdate(iType ip.IFlag, rType dns.RType) error {
	var local, remote, record string
	var err error

	// get local address
	if local, err = c.ip.Get(iType); err != nil {
		return err
	}
	c.Log.Info("local address", "iType", iType, "local", local)

	// get remote address
	if remote, err = c.dns.Get(c.config.zone, c.config.name, rType); err != nil {
		return err
	}
	c.Log.Info("remote address", "iType", iType, "remote", remote)

	// only update if needed
	if local == remote {
		c.Log.Debug("skipping update - no change detected")
		return nil
	}

	// attempt to update remote record
	if record, err = c.dns.Upsert(c.config.zone, c.config.name, local, rType); err != nil {
		return err
	}
	c.Log.Info("updated remote", "record", record, "rType", rType, "data", local)

	return nil
}
