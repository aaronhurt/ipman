package check

import (
	stderrors "errors"
	"flag"
	"fmt"
	"os"

	"github.com/leprechau/ipman/internal"
	internalerrors "github.com/leprechau/ipman/internal/errors"
)

// setupFlags initializes the instance configuration
func (c *Command) setupFlags(args []string) (bool, error) {
	var cmdFlags *flag.FlagSet // instance flagset
	var err error

	// init config if needed
	if c.config == nil {
		c.config = new(config)
	}

	// init flagset
	cmdFlags = flag.NewFlagSet("check", flag.ContinueOnError)
	cmdFlags.SetOutput(os.Stdout)
	cmdFlags.Usage = func() { _, _ = os.Stdout.WriteString(c.Help()) }

	// declare flags
	cmdFlags.BoolVar(&c.config.v4, "4", false,
		"Check IPv4")
	cmdFlags.BoolVar(&c.config.v6, "6", false,
		"Check IPv6")
	cmdFlags.StringVar(&c.config.ipbe, "ipbe", "ipify",
		"IP lookup backend")

	// parse flags and catch help
	if err = cmdFlags.Parse(args); err != nil {
		if stderrors.Is(err, flag.ErrHelp) {
			return true, nil
		}
		return false, fmt.Errorf("parse check flags: %w", err)
	}

	// check for remaining garbage
	if cmdFlags.NArg() > 0 {
		return false, internalerrors.ErrUnknownArg
	}

	// default to v4 if not specified
	if !c.config.v4 && !c.config.v6 {
		c.config.v4 = true
	}

	// init ip backend
	if c.ip, err = internal.GetIPBackend(c.config.ipbe); err != nil {
		return false, fmt.Errorf("init ip backend: %w", err)
	}

	return false, nil
}
