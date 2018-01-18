package godaddy

import (
	// core
	"net/http"
	"os"

	// Hashi's clean http client
	"github.com/hashicorp/go-cleanhttp"

	// get domain from hostname
	"golang.org/x/net/publicsuffix"

	// ENA's simple restclient
	"github.com/myENA/restclient"
)

// Config contains backend configuration
type Config struct {
	defaultURL        string
	accessKey         string
	secretKey         string
	defaultDomainName string
	defaultRecordName string
	defaultRecordTTL  int
	client            *restclient.Client
}

// DefaultConfig returns the default backend configuration
func DefaultConfig() (*Config, error) {
	var config *Config
	var host string
	var err error

	// init base config
	config = &Config{
		defaultURL:        "https://api.godaddy.com/v1",
		defaultRecordName: "@",
		defaultRecordTTL:  600,
	}

	// init client
	config.client = &restclient.Client{
		Client: cleanhttp.DefaultClient(),
	}

	// set error handler
	config.client.ErrorResponseCallback = func(resp *http.Response) error {
		return parseError(resp)
	}

	// attempt to get hostname
	if host, err = os.Hostname(); err != nil {
		return config, err
	}

	// attempt to get domain from host
	config.defaultDomainName, err = publicsuffix.EffectiveTLDPlusOne(host)
	if err != nil {
		return config, err
	}

	// all good
	return config, nil
}

// AccessKey allows setting the API access key
func (c *Config) AccessKey(key string) {
	c.accessKey = key
}

// SecretKey allows setting the API secret key
func (c *Config) SecretKey(key string) {
	c.secretKey = key
}
