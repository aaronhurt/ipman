// Package cloudflare provides a DNS backend via the Cloudflare DNS service
package cloudflare

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/leprechau/ipman/internal/dns"
	"github.com/leprechau/ipman/internal/errors"
)

// Get a record by name
func (c *Config) Get(zone, name string, typ dns.RType) (string, error) {
	var err error
	var r *resty.Response

	// cloudflare punycode mapping seems broken - lookup the zone apex instead
	if name == "@" {
		if r, err = c.client.R().
			SetAuthToken(c.apiToken).
			SetPathParam("zone", zone).
			SetResult(&ZoneResponse{}).
			SetError(&DNSErrorResponse{}).
			SetDebug(true).
			Get("/zones/{zone}"); err != nil {
			return "", err
		}

		if r.IsError() {
			return "", formatDNSError(r)
		}

		res, ok := r.Result().(*ZoneResponse)
		if !ok {
			return "", errors.ErrUnexpectedResponse
		}

		name = res.Result.Name
	}

	if r, err = c.client.R().
		SetAuthToken(c.apiToken).
		SetQueryParam("name", name).
		SetQueryParam("type", typ.String()).
		SetPathParam("zone", zone).
		SetResult(&DNSResponse{}).
		SetError(&DNSErrorResponse{}).
		SetDebug(true).
		Get("/zones/{zone}/dns_records"); err != nil {
		return "", err
	}

	if r.IsError() {
		return "", formatDNSError(r)
	}

	res, ok := r.Result().(*DNSResponse)
	if !ok {
		return "", errors.ErrUnexpectedResponse
	}

	if res.ResultInfo.Count < 1 {
		return "", nil
	}

	c.recordID = res.Result[0].ID
	return res.Result[0].Content, nil
}

// Upsert a record by name
func (c *Config) Upsert(zone, name, data string, typ dns.RType) (string, error) {
	r, err := c.client.R().
		SetAuthToken(c.apiToken).
		SetPathParam("zone", zone).
		SetPathParam("record", c.recordID).
		SetBody(&DNSRecord{
			Content: data,
			Name:    name,
			Type:    typ.String(),
			ID:      c.recordID,
			TTL:     c.recordTTL,
		}).
		SetResult(&DNSUpdateResponse{}).
		SetError(&DNSErrorResponse{}).
		Patch("/zones/{zone}/dns_records/{record}")

	if err != nil {
		return "", err
	}

	if r.IsError() {
		return "", formatDNSError(r)
	}

	res, ok := r.Result().(*DNSUpdateResponse)
	if !ok {
		return "", errors.ErrUnexpectedResponse
	}

	return res.Result.Name, nil
}

// RecordTTL returns the default record ttl
func (c *Config) RecordTTL() int {
	return c.recordTTL
}

// formatDNSError formats an upstream DNS Error response
func formatDNSError(r *resty.Response) error {
	res, ok := r.Error().(*DNSErrorResponse)
	if !ok {
		return errors.ErrUnexpectedResponse
	}
	return fmt.Errorf("status: %d, success: %t, errors: %+v, messages: %+v",
		r.StatusCode(), res.Success, res.Errors, res.Messages)
}
