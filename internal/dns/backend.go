// Package dns defines the DNS service backend interface
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
	Upsert(domain, name, data string, typ RType) (string, error)

	// RecordTTL returns the default record ttl
	RecordTTL() int

	// SetAccessKey sets an API access key
	SetAccessKey(string)

	// SetSecretKey sets an API token or secret key
	SetSecretKey(string)
}
