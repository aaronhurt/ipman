package check

import (
	"flag"
	"fmt"
	"os"

	// common errors
	ce "github.com/leprechau/ipman/common/errors"

	// provider backends
	"github.com/leprechau/ipman/common/ip/myexternal"
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
	cmdFlags.Usage = func() { fmt.Fprint(os.Stdout, c.Help()); os.Exit(0) }

	// declare flags
	cmdFlags.BoolVar(&c.config.v4, "4", false,
		"Check IPv4")
	cmdFlags.BoolVar(&c.config.v6, "6", false,
		"Check IPv6")

	// parse flags and ignore error
	if err = cmdFlags.Parse(args); err != nil {
		return nil
	}

	// check for remaining garbage
	if cmdFlags.NArg() > 0 {
		return ce.ErrUnknownArg
	}

	// default to v4 if not specified
	if !c.config.v4 && !c.config.v6 {
		c.config.v4 = true
	}

	// init ip backend (currently only one)
	if c.ip, err = myexternal.DefaultConfig(); err != nil {
		return err
	}

	// always okay
	return nil
}
