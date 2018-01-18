package ip

// IFlag represents an address family inet flag
type IFlag uint8

// Inet copies the AF_INET definition
const Inet IFlag = 0x2

// Inet6 copies the AF_INET definition
const Inet6 IFlag = 0xa

// String implements the stringer interface
func (i IFlag) String() string {
	if i == Inet {
		return "IPv4"
	}
	if i == Inet6 {
		return "IPv6"
	}
	// default
	return "Unknown"
}

// Backend defines the IP lookup backend interface
type Backend interface {
	// Get the host IP address for the given protocol
	Get(proto IFlag) (string, error)
}

// Default represents the default IP lookup backend
var Default Backend
