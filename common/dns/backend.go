package dns

// RType represents a dns record type
type RType string

// A record constant
const A RType = "A"

// AAAA record constant flag
const AAAA RType = "AAAA"

// String implements the stringer interface
func (r RType) String() string {
	return string(r)
}

// Backend defines the DNS registry backend interface
type Backend interface {
	// Get a record via a DNS backend
	Get(domain, name string, typ RType) (string, error)

	// Upsert a record into a DNS backend
	Upsert(domain, name, data string, typ RType) error

	// DefaultDomainName returns the default domain name
	DefaultDomainName() string

	// DefaultRecordName returns the default record name
	DefaultRecordName() string

	// DefaultRecordTTL returns the default record ttl
	DefaultRecordTTL() int

	// AccessKey provides a method to set an API access key
	AccessKey(string)

	// SecretKey provides a method to set an API secret key
	SecretKey(string)
}
