package common

import (
	"errors"
)

// ErrUnknownArg is returned when non-flag arguments are present after the command
var ErrUnknownArg = errors.New("unknown non-flag argument(s) present after command")

// ErrUnknownIPBackend is returned when an unknown IP backend is specified on the cli
var ErrUnknownIPBackend = errors.New("unknown IP backend specified for command")

// ErrUnknownDNSBackend is returned when an unknown DNS backend is specified on the cli
var ErrUnknownDNSBackend = errors.New("unknown DNS backend specified for command")
