package update

import (
	"flag"
	"fmt"
	"os"

	// errors package
	ce "github.com/leprechau/ipman/common/errors"

	// provider backends
	"github.com/leprechau/ipman/common/dns/godaddy"
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
	cmdFlags = flag.NewFlagSet("update", flag.ContinueOnError)
	cmdFlags.Usage = func() { fmt.Fprint(os.Stdout, c.Help()); os.Exit(0) }

	// declare flags
	cmdFlags.BoolVar(&c.config.v4, "4", false,
		"Update IPv4 (A) record")
	cmdFlags.BoolVar(&c.config.v6, "6", false,
		"Update IPv6 (AAAA) record")
	cmdFlags.StringVar(&c.config.key, "key", "",
		"API access key")
	cmdFlags.StringVar(&c.config.secret, "secret", "",
		"API secret key")
	cmdFlags.StringVar(&c.config.domain, "domain", "",
		"Domain name")
	cmdFlags.StringVar(&c.config.name, "record", "@",
		"Record name")
	cmdFlags.IntVar(&c.config.ttl, "ttl", 600,
		"TTL value")

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

	// attempt to populate key from environment
	if c.config.key == "" {
		c.config.key = os.Getenv("IPMAN_DNS_KEY")
	}

	// attempt to populate secret from environment
	if c.config.secret == "" {
		c.config.secret = os.Getenv("IPMAN_DNS_SECRET")
	}

	// init ip backend (currently only one)
	if c.ip, err = myexternal.DefaultConfig(); err != nil {
		return err
	}

	// init dns backend (currently only one)
	if c.dns, err = godaddy.DefaultConfig(); err != nil {
		return err
	}

	// set backend access key if needed
	if c.config.key != "" {
		c.dns.AccessKey(c.config.key)
	}

	// set backend secret key if needed
	if c.config.secret != "" {
		c.dns.SecretKey(c.config.secret)
	}

	// check domain
	if c.config.domain == "" {
		c.config.domain = c.dns.DefaultDomainName()
	}

	// check name
	if c.config.name == "" {
		c.config.name = c.dns.DefaultRecordName()
	}

	// check ttl
	if c.config.ttl == 0 {
		c.config.ttl = c.dns.DefaultRecordTTL()
	}

	// always okay
	return nil
}
