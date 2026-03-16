package check

import (
	"strings"
	"testing"

	internalerrors "github.com/leprechau/ipman/internal/errors"
)

func TestSetupFlagsDefaultsToIPv4(t *testing.T) {
	t.Parallel()

	cmd := &Command{Self: "ipman"}

	if err := cmd.setupFlags(nil); err != nil {
		t.Fatalf("setupFlags() error = %v", err)
	}

	if !cmd.config.v4 {
		t.Fatal("setupFlags() did not default to IPv4")
	}

	if cmd.config.v6 {
		t.Fatal("setupFlags() unexpectedly enabled IPv6")
	}

	if cmd.ip == nil {
		t.Fatal("setupFlags() did not initialize an IP backend")
	}
}

func TestSetupFlagsReturnsParseError(t *testing.T) {
	t.Parallel()

	cmd := &Command{Self: "ipman"}

	err := cmd.setupFlags([]string{"-unknown"})
	if err == nil {
		t.Fatal("setupFlags() error = nil, want parse failure")
	}

	if !strings.Contains(err.Error(), "parse check flags") {
		t.Fatalf("setupFlags() error = %v, want contextual parse error", err)
	}
}

func TestSetupFlagsRejectsTrailingArgs(t *testing.T) {
	t.Parallel()

	cmd := &Command{Self: "ipman"}

	err := cmd.setupFlags([]string{"extra"})
	if err == nil {
		t.Fatal("setupFlags() error = nil, want trailing arg failure")
	}

	if err != internalerrors.ErrUnknownArg {
		t.Fatalf("setupFlags() error = %v, want %v", err, internalerrors.ErrUnknownArg)
	}
}
