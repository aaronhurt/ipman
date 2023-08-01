package common

import (
	"errors"
)

// ErrUnknownArg is returned when non-flag arguments are present after the command
var ErrUnknownArg = errors.New("Unknown non-flag argument(s) present after command")
var ErrUnknownIPBackend = errors.New("Unknown IP backend specified for command")
var ErrUnknownDNSBackend = errors.New("Unknown DNS backend specified for command")
