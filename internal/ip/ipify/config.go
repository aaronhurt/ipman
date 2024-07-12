package ipify

import (
	"github.com/go-resty/resty/v2"
)

// Config contains backend configuration
type Config struct {
	v4URL  string
	v6URL  string
	client *resty.Client
}

// DefaultConfig returns the default backend configuration
func DefaultConfig() (*Config, error) {
	return &Config{
		v4URL: "https://api.ipify.org",
		v6URL: "https://api6.ipify.org",
		client: resty.New().
			SetHeader("Accept", "application/json"),
	}, nil
}
