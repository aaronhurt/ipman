package check

import (
	// core
	"fmt"

	// backend definition
	"github.com/leprechau/ipman/common/ip"
)

// checkIP
func (c *Command) checkIP() error {
	var ips string
	var err error

	// get local v4
	if c.config.v4 {
		if ips, err = c.ip.Get(ip.Inet); err != nil {
			return err
		}
		fmt.Println(ips)
	}

	// get local v6
	if c.config.v6 {
		if ips, err = c.ip.Get(ip.Inet6); err != nil {
			return err
		}
		fmt.Println(ips)
	}

	// all okay
	return nil
}
