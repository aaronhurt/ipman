package godaddy

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	// backend definition
	"github.com/leprechau/ipman/common/dns"
)

// Get a record via the GoDaddy API
func (c *Config) Get(domain, name string, typ dns.RType) (string, error) {
	var ctx = context.Background()
	var response = make(DNSRecords, 0)
	var u *url.URL
	var path string
	var err error

	// parse url and check error
	if u, err = url.Parse(c.defaultURL); err != nil {
		return "", err
	}

	// build path
	path = fmt.Sprintf("/domains/%s/records/%s/%s",
		domain, typ, name)

	// ensure auth header is set and current
	c.setAuthorizationHeader()

	// execute client call
	if err = c.client.Get(ctx, u, path, nil, &response); err != nil {
		return "", err
	}

	// check for empty response - no upstream records
	if len(response) < 1 {
		return "", nil
	}

	// return the payload
	return strings.ToLower(response[0].Data), nil
}

// Upsert a record via the GoDaddy API
func (c *Config) Upsert(domain, name, data string, typ dns.RType) error {
	var ctx = context.Background()
	var request = make(DNSRecords, 1)
	var u *url.URL
	var path string
	var err error

	// parse url and check error
	if u, err = url.Parse(c.defaultURL); err != nil {
		return err
	}

	// build path
	path = fmt.Sprintf("/domains/%s/records/%s/%s",
		domain, typ, name)

	// populate request
	request[0].Type = typ.String()
	request[0].Name = name
	request[0].Data = data

	// ensure auth header is set and current
	c.setAuthorizationHeader()

	// execute client call and return
	return c.client.Put(ctx, u, path, nil, request, nil)
}

// DefaultDomainName returns the default domain name
func (c *Config) DefaultDomainName() string {
	return c.defaultDomainName
}

// DefaultRecordName returns the default record name
func (c *Config) DefaultRecordName() string {
	return c.defaultRecordName
}

// DefaultRecordTTL returns the default record ttl
func (c *Config) DefaultRecordTTL() int {
	return c.defaultRecordTTL
}

// authorizationHeader builds and sets the authorization header in the client
func (c *Config) setAuthorizationHeader() {
	c.client.FixupCallback = func(req *http.Request) error {
		req.Header.Set("Authorization",
			fmt.Sprintf("sso-key %s:%s", c.accessKey, c.secretKey))
		return nil
	}
}

// parseError handles 400+ error codes returned from the API endpoint
func parseError(resp *http.Response) error {
	var obj interface{} // generic object definition
	var err error

	// close body when done
	defer resp.Body.Close()

	// check response code and set object
	if resp.StatusCode == http.StatusTooManyRequests {
		obj = new(DNSErrorLimit)
	} else {
		obj = new(DNSError)
	}

	// attempt process error body
	if err = json.NewDecoder(resp.Body).Decode(obj); err != nil {
		return fmt.Errorf("Unable to process error from %d error response: %s",
			resp.StatusCode, err.Error())
	}

	// check for fields
	if len(obj.(*DNSError).Fields) > 0 {
		// return error with fields
		return fmt.Errorf("[%d/%s] %s -> %v",
			resp.StatusCode, obj.(*DNSError).Code,
			obj.(*DNSError).Message, obj.(*DNSError).Fields)
	}

	// return error without fields
	return fmt.Errorf("[%d/%s] %s",
		resp.StatusCode, obj.(*DNSError).Code,
		obj.(*DNSError).Message)
}
