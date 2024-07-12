package check

import (
	"flag"
	"fmt"
	"os"

	"github.com/leprechau/ipman/internal"
	"github.com/leprechau/ipman/internal/errors"
)

// setupFlags initializes the instance configuration
func (c *Command) setupFlags(args []string) error {
	var cmdFlags *flag.FlagSet // instance flagset
	var err error

	// init config if needed
	if c.config == nil {
		c.config = new(config)
	}

	// init flagset
	cmdFlags = flag.NewFlagSet("check", flag.ContinueOnError)
	cmdFlags.Usage = func() { _, _ = fmt.Fprint(os.Stdout, c.Help()); os.Exit(0) }

	// declare flags
	cmdFlags.BoolVar(&c.config.v4, "4", false,
		"Check IPv4")
	cmdFlags.BoolVar(&c.config.v6, "6", false,
		"Check IPv6")
	cmdFlags.StringVar(&c.config.ipbe, "ipbe", "ipify",
		"IP lookup backend")

	// parse flags and ignore error
	if err = cmdFlags.Parse(args); err != nil {
		return nil
	}

	// check for remaining garbage
	if cmdFlags.NArg() > 0 {
		return errors.ErrUnknownArg
	}

	// default to v4 if not specified
	if !c.config.v4 && !c.config.v6 {
		c.config.v4 = true
	}

	// init ip backend
	if c.ip, err = internal.GetIPBackend(c.config.ipbe); err != nil {
		return err
	}

	// all okay
	return nil
}
