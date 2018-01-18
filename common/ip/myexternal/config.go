package myexternal

import (
	// ENA's simple restclient
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
		v4URL: "https://ipv4.myexternalip.com",
		v6URL: "https://ipv6.myexternalip.com", // has invalid cert :/
	}

	// init global client
	config.client, err = restclient.NewClient(
		&restclient.ClientConfig{
			InsecureSkipVerify: true, // ipv6.myexternalip.com
		},
		nil,
	)

	// return error
	return config, err
}
