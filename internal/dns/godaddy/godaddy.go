// Package godaddy provides a DNS backend via the GoDaddy service
package godaddy

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/leprechau/ipman/internal/dns"
	"github.com/leprechau/ipman/internal/errors"
)

// Get a record via by name
func (c *Config) Get(zone, name string, typ dns.RType) (string, error) {
	r, err := c.client.R().
		SetPathParam("zone", zone).
		SetPathParam("name", name).
		SetPathParam("type", typ.String()).
		SetResult(DNSRecords{}).
		SetError(&DNSError{}).
		SetAuthScheme("sso-key").
		SetAuthToken(c.accessKey + ":" + c.secretKey).
		Get("/domains/{zone}/records/{type}/{name}")

	if err != nil {
		return "", err
	}

	if r.IsError() {
		return "", formatDNSError(r)
	}

	res, ok := r.Result().(DNSRecords)
	if !ok {
		return "", errors.ErrUnexpectedResponse
	}

	// check for empty response - no upstream records
	if len(res) < 1 {
		return "", nil
	}

	// return the payload
	return strings.ToLower(res[0].Data), nil
}

// Upsert a record by name
func (c *Config) Upsert(zone, name, data string, typ dns.RType) (string, error) {
	r, err := c.client.R().
		SetPathParam("zone", zone).
		SetQueryParam("name", name).
		SetQueryParam("type", typ.String()).
		SetBody(DNSRecords{
			{
				Type: typ.String(),
				Name: name,
				Data: data,
			},
		}).
		SetError(&DNSError{}).
		SetAuthScheme("sso-key").
		SetAuthToken(c.accessKey + ":" + c.secretKey).
		Put("/domains/{zone}/records/{type}/{name}")

	if err != nil {
		return "", err
	}

	if r.IsError() {
		return "", formatDNSError(r)
	}

	return data, nil
}

// RecordTTL returns the default record ttl
func (c *Config) RecordTTL() int {
	return c.recordTTL
}

// formatDNSError formats an upstream DNS Error response
func formatDNSError(r *resty.Response) error {
	res, ok := r.Error().(*DNSError)
	if !ok {
		return errors.ErrUnexpectedResponse
	}
	// return error without fields
	return fmt.Errorf("status: %d, code: %s, message: %s, fields: %+v",
		r.StatusCode(), res.Code, res.Message, res.Fields)
}
