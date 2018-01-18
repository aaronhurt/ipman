package godaddy

// DNSRecord represents an element of the GoDaddy model
// https://developer.godaddy.com/doc#!/_v1_domains/recordReplace/ArrayOfDNSRecord
type DNSRecord struct {
	Type     string  `json:"type" validate:"required,eq=A|eq=AAAA|eq=CNAME|eq=MX|eq=NS|eq=SOA|eq=SRV|eq=TXT"`
	Name     string  `json:"name" validate:"required,min=1,max=255"`
	Data     string  `json:"data" validate:"required,min=1,max=255"`
	Priority *int    `json:"priority,omitempty" validate:"omitempty,gte=1"`
	TTL      *int    `json:"ttl,omitempty" validate:"omitempty,gte=1"`
	Service  *string `json:"service,omitempty" validate:"omitempty,min=1"`
	Protocol *string `json:"protocol,omitempty" validate:"omitempty,min=1"`
	Port     *int    `json:"port,omitempty" validate:"omitempty,min=1,max=65535"`
	Weight   *int    `json:"weight,omitempty" validate:"omitempty,gte=1"`
}

// DNSRecords represents the GoDaddy model
// https://developer.godaddy.com/doc#!/_v1_domains/recordReplace/ArrayOfDNSRecord
type DNSRecords []DNSRecord

// DNSError represents the GoDaddy model shared across the domain APIs
// https://developer.godaddy.com/doc#!/_v1_domains/list/Error
type DNSError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Fields  []struct {
		Path        string `json:"path"`
		PathRelated string `json:"pathRelated"`
		Code        string `json:"code"`
		Message     string `json:"message"`
	} `json:"fields"`
}

// DNSErrorLimit represents the GoDaddy model shared across domain APIs
// https://developer.godaddy.com/doc#!/_v1_domains/list/ErrorLimit
type DNSErrorLimit struct {
	RetryAfterSec int `json:"retryAfterSec"`
	DNSError
}
