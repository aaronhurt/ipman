package godaddy

import (
	"github.com/go-resty/resty/v2"
)

// Config contains backend configuration
type Config struct {
	accessKey string
	secretKey string
	recordTTL int
	client    *resty.Client
}

// DefaultConfig returns the default backend configuration
func DefaultConfig() (*Config, error) {
	return &Config{
		recordTTL: 600,
		client: resty.New().
			SetHeader("Accept", "application/json").
			SetBaseURL("https://api.godaddy.com/v1"),
	}, nil
}

// SetAccessKey sets the API access key
func (c *Config) SetAccessKey(key string) {
	c.accessKey = key
}

// SetSecretKey sets the API secret key
func (c *Config) SetSecretKey(key string) {
	c.secretKey = key
}
