package ipify

import (
	"github.com/myENA/restclient"
)

// Config contains backend configuration
type Config struct {
	v4URL  string
	v6URL  string
	client *restclient.Client
}

// DefaultConfig returns the default backend configuration
func DefaultConfig() (*Config, error) {
	var config *Config
	var err error

	config = &Config{
		v4URL: "https://api.ipify.org",
		v6URL: "https://api6.ipify.org",
	}

	// init global client
	config.client, err = restclient.NewClient(&restclient.ClientConfig{}, nil)

	// return error
	return config, err
}
