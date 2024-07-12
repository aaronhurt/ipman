package godaddy

// DNSRecord represents an element of the GoDaddy model
// https://developer.godaddy.com/doc/endpoint/domains#/v1/recordGet
type DNSRecord struct {
	Type     string  `json:"type"`
	Name     string  `json:"name"`
	Data     string  `json:"data"`
	Priority *int    `json:"priority,omitempty"`
	TTL      *int    `json:"ttl,omitempty"`
	Service  *string `json:"service,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
	Port     *int    `json:"port,omitempty"`
	Weight   *int    `json:"weight,omitempty"`
}

// DNSRecords represents the GoDaddy model
// https://developer.godaddy.com/doc/endpoint/domains#/v1/recordGet
type DNSRecords []DNSRecord

// DNSError represents the GoDaddy model shared across the domain APIs
// https://developer.godaddy.com/doc/endpoint/domains#/v1/recordGet
type DNSError struct {
	Code   string `json:"code"`
	Fields []struct {
		Code        string `json:"code"`
		Message     string `json:"message"`
		Path        string `json:"path"`
		PathRelated string `json:"pathRelated"`
	} `json:"fields"`
	Message       string `json:"message"`
	RetryAfterSec int    `json:"retryAfterSec"`
}
