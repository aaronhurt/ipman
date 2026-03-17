package ipify

import (
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/leprechau/ipman/internal/ip"
)

const requestTimeout = 30 * time.Second

// Config contains backend configuration
type Config struct {
	v4URL  string
	v6URL  string
	client *resty.Client
}

var _ ip.Backend = (*Config)(nil)

// DefaultConfig returns the default backend configuration
func DefaultConfig() *Config {
	return &Config{
		v4URL: "https://api.ipify.org",
		v6URL: "https://api6.ipify.org",
		client: resty.New().
			SetHeader("Accept", "application/json").
			SetTimeout(requestTimeout),
	}
}
