package update

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
	cmdFlags = flag.NewFlagSet("update", flag.ContinueOnError)
	cmdFlags.Usage = func() { _, _ = fmt.Fprint(os.Stdout, c.Help()); os.Exit(0) }

	// declare flags
	cmdFlags.BoolVar(&c.config.v4, "4", false,
		"Update IPv4 (A) record")
	cmdFlags.BoolVar(&c.config.v6, "6", false,
		"Update IPv6 (AAAA) record")
	cmdFlags.StringVar(&c.config.key, "key", "",
		"API access key (optional)")
	cmdFlags.StringVar(&c.config.secret, "secret", "",
		"API secret key or token")
	cmdFlags.StringVar(&c.config.zone, "zone", "",
		"SetZone ID or domain name")
	cmdFlags.StringVar(&c.config.name, "name", "@",
		"Record name")
	cmdFlags.IntVar(&c.config.ttl, "ttl", 600,
		"TTL value")
	cmdFlags.StringVar(&c.config.ipbe, "ipbe", "ipify",
		"IP lookup backend")
	cmdFlags.StringVar(&c.config.dnsbe, "dnsbe", "cloudflare",
		"DNS update backend")

	// parse flags and ignore error
	if err = cmdFlags.Parse(args); err != nil {
		return nil
	}

	// check for remaining garbage
	if cmdFlags.NArg() > 0 {
		return errors.ErrUnknownArg
	}

	// check zone and attempt to get from environment
	if c.config.zone == "" {
		if c.config.zone = os.Getenv("IPMAN_DNS_ZONE"); c.config.zone == "" {
			return errors.ErrMissingZone
		}
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

	// init ip backend
	if c.ip, err = internal.GetIPBackend(c.config.ipbe); err != nil {
		return err
	}

	// init dns backend
	if c.dns, err = internal.GetDNSBackend(c.config.dnsbe); err != nil {
		return err
	}

	// set backend access key if needed
	if c.config.key != "" {
		c.dns.SetAccessKey(c.config.key)
	}

	// set backend secret key if needed
	if c.config.secret != "" {
		c.dns.SetSecretKey(c.config.secret)
	}

	// check ttl
	if c.config.ttl == 0 {
		c.config.ttl = c.dns.RecordTTL()
	}

	// always okay
	return nil
}
