package update

import (
	"fmt"
	stdLog "log"

	// backend definitions
	"github.com/leprechau/ipman/common/dns"
	"github.com/leprechau/ipman/common/ip"
)

// command configuration
type config struct {
	v4     bool
	v6     bool
	key    string
	secret string
	domain string
	name   string
	ttl    int
	ipbe   string
	dnsbe  string
}

// Command is a Command implementation for the update operation
type Command struct {
	Self   string
	Log    *stdLog.Logger
	config *config
	dns    dns.Backend
	ip     ip.Backend
}

// Run is a function to run the command
func (c *Command) Run(args []string) int {
	var err error

	// init flags
	if err = c.setupFlags(args); err != nil {
		c.Log.Printf("[Error] Failed to init flags: %s", err.Error())
		return 1
	}

	// attempt to update p
	if err = c.updateIP(); err != nil {
		c.Log.Printf("[Error] Failed to update dns record: %s", err.Error())
		return 1
	}

	// exit clean
	return 0
}

// Synopsis shows the command summary
func (c *Command) Synopsis() string {
	return "Update DNS registry with external ip address of local machine."
}

// Help shows the detailed command options
func (c *Command) Help() string {
	return fmt.Sprintf(`
Usage: %s cmd [options]

	Update DNS registry with external ip address of local machine.

Options:

	-4        Update external IPv4 address if available.
	-6        Update external IPv6 address if available.
	-key      The DNS API access key.
	-secret   The DNS API access secret.
	-domain   The DNS domain name.                   (default: local domain)
	-name     The DNS record name.                   (default: @)
	-ttl      The DNS record ttl in seconds.         (default: 600)
	-ipbe     IP lookup backend (local or ipify).    (default: ipify)
	-dnsbe    DNS update backend (godaddy or xxx)    (default: godaddy)
`, c.Self)
}
