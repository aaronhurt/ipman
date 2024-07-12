// Package errors contains common errors used within the application
package errors

import (
	"errors"
)

// ErrUnknownArg is returned when non-flag arguments are present after the command
var ErrUnknownArg = errors.New("unknown non-flag argument(s) present after command")

// ErrUnknownIPBackend is returned when an unknown IP backend is specified on the cli
var ErrUnknownIPBackend = errors.New("unknown IP backend specified for command")

// ErrUnknownDNSBackend is returned when an unknown DNS backend is specified on the cli
var ErrUnknownDNSBackend = errors.New("unknown DNS backend specified for command")

// ErrMissingZone is returned when no zone ID or domain name is specified on the cli
var ErrMissingZone = errors.New("missing zone or domain name")

// ErrUnexpectedResponse is returned when the upstream API returns an unexpected response
var ErrUnexpectedResponse = errors.New("received unexpected response from upstream service")
