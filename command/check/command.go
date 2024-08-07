package check

import (
	"fmt"
	"log/slog"

	"github.com/leprechau/ipman/internal/ip"
)

// command configuration
type config struct {
	v4   bool
	v6   bool
	ipbe string
}

// Command is a Command implementation for the check operation
type Command struct {
	Self   string
	Log    *slog.Logger
	config *config
	ip     ip.Backend
}

// Run is a function to run the command
func (c *Command) Run(args []string) int {
	var err error

	// init flags
	if err = c.setupFlags(args); err != nil {
		c.Log.Error("failed to init flags", "err", err)
		return 1
	}

	// check ip
	if err = c.checkIP(); err != nil {
		c.Log.Error("failed to check addresses", "err", err)
		return 1
	}

	// exit clean
	return 0
}

// Synopsis shows the command summary
func (c *Command) Synopsis() string {
	return "Return current external ip address of local machine."
}

// Help shows the detailed command options
func (c *Command) Help() string {
	return fmt.Sprintf(`
Usage: %s cmd [options]

	Return current external ip address of local machine.

Options:

	-4        Get external IPv4 address if available.
	-6        Get external IPv6 address if available.
	-ipbe     IP lookup backend (ipify or local).          (default: ipify)
`, c.Self)
}
