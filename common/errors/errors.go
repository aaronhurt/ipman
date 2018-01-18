package errors

import (
	"errors"
)

// ErrUnknownArg is returned when non-flag arguments are present after the command
var ErrUnknownArg = errors.New("Unknown non-flag argument(s) present after command")
