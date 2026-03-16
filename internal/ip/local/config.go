package local

import "github.com/leprechau/ipman/internal/ip"

// Config contains backend configuration
type Config struct {
	v4Host string
	v6Host string
}

var _ ip.Backend = (*Config)(nil)

// DefaultConfig returns the default backend configuration
func DefaultConfig() *Config {
	// google public dns - the connection does not need to succeed
	return &Config{
		v4Host: "8.8.8.8:53",
		v6Host: "[2001:4860:4860::8888]:53",
	}
}
