// Package ipify provides an IP lookup backend via the ipify.org service
package ipify

import (
	"strings"

	"github.com/leprechau/ipman/internal/errors"
	"github.com/leprechau/ipman/internal/ip"
)

// Get looks up client address via ipify
func (c *Config) Get(proto ip.IFlag) (string, error) {
	url := c.v4URL
	if proto == ip.Inet6 {
		url = c.v6URL
	}

	r, err := c.client.R().
		SetQueryParam("format", "json").
		SetResult(&Response{}).
		Get(url)

	if err != nil {
		return "", err
	}

	res, ok := r.Result().(*Response)
	if !ok {
		return "", errors.ErrUnexpectedResponse
	}

	// return the ip payload
	return strings.ToLower(res.IP), nil
}
