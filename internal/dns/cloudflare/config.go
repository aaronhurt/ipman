package cloudflare

import (
	"github.com/go-resty/resty/v2"
	"github.com/leprechau/ipman/internal/dns"
)

// Config contains backend configuration
type Config struct {
	apiToken  string
	recordID  string // populated by Get and only valid till the next call to Get
	recordTTL int
	client    *resty.Client
}

var _ dns.Backend = (*Config)(nil)

// DefaultConfig returns the default backend configuration
func DefaultConfig() *Config {
	return &Config{
		recordTTL: 600,
		client: resty.New().
			SetHeader("Accept", "application/json").
			SetBaseURL("https://api.cloudflare.com/client/v4"),
	}
}

// SetAccessKey sets the API access key
func (c *Config) SetAccessKey(string) {
	// cloudflare uses a single API Token
}

// SetSecretKey sets the API access token
func (c *Config) SetSecretKey(s string) {
	c.apiToken = s
}
