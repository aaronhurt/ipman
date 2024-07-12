package cloudflare

import (
	"time"
)

// DNSResponse represents a DNS query response
// https://developers.cloudflare.com/api/operations/dns-records-for-a-zone-list-dns-records
type DNSResponse struct {
	Success    bool              `json:"success"`
	Errors     []ResponseError   `json:"errors"`
	Messages   []ResponseMessage `json:"messages"`
	Result     []DNSRecord       `json:"result"`
	ResultInfo ResponseInfo      `json:"result_info"`
}

// DNSUpdateResponse represents a DNS update response
// https://developers.cloudflare.com/api/operations/dns-records-for-a-zone-patch-dns-record
type DNSUpdateResponse struct {
	Success  bool              `json:"success"`
	Errors   []ResponseError   `json:"errors"`
	Messages []ResponseMessage `json:"messages"`
	Result   DNSRecord         `json:"result"`
}

// DNSRecord represents a single result within a DNSResponse
type DNSRecord struct {
	ID         string     `json:"id"`
	Content    string     `json:"content"`
	Name       string     `json:"name"`
	Type       string     `json:"type"`
	TTL        int        `json:"ttl"`
	ZoneID     *string    `json:"zone_id,omitempty"`
	ZoneName   *string    `json:"zone_name,omitempty"`
	Proxiable  *bool      `json:"proxiable,omitempty"`
	Proxied    *bool      `json:"proxied,omitempty"`
	Locked     *bool      `json:"locked,omitempty"`
	Comment    *string    `json:"comment,omitempty"`
	Tags       *[]string  `json:"tags,omitempty"`
	CreatedOn  *time.Time `json:"created_on,omitempty"`
	ModifiedOn *time.Time `json:"modified_on,omitempty"`
}

// ZoneResponse represents a DNS zone response
// https://developers.cloudflare.com/api/operations/zones-0-get
type ZoneResponse struct {
	Success  bool              `json:"success"`
	Errors   []ResponseError   `json:"errors"`
	Messages []ResponseMessage `json:"messages"`
	Result   ZoneResult        `json:"result"`
}

// ZoneResult represents a ZoneResponse result element
type ZoneResult struct {
	ID                  string    `json:"id"`
	Name                string    `json:"name"`
	Status              string    `json:"status"`
	Paused              bool      `json:"paused"`
	Type                string    `json:"type"`
	DevelopmentMode     int       `json:"development_mode"`
	NameServers         []string  `json:"name_servers"`
	OriginalNameServers []string  `json:"original_name_servers"`
	OriginalRegistrar   string    `json:"original_registrar"`
	OriginalDnshost     string    `json:"original_dnshost"`
	ModifiedOn          time.Time `json:"modified_on"`
	CreatedOn           time.Time `json:"created_on"`
	ActivatedOn         time.Time `json:"activated_on"`
}

// DNSErrorResponse represents a DNS query response error
type DNSErrorResponse struct {
	Success  bool              `json:"success"`
	Errors   []ResponseError   `json:"errors"`
	Messages []ResponseMessage `json:"messages"`
}

// ResponseError represents the common errors object used in API responses
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ResponseMessage represents the common messages object used in API responses
type ResponseMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ResponseInfo represents the common result_info object used in API responses
type ResponseInfo struct {
	Count      int `json:"count"`
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	TotalCount int `json:"total_count"`
}
