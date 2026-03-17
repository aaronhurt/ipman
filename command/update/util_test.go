package update

import (
	"strings"
	"testing"

	internalerrors "github.com/leprechau/ipman/internal/errors"
)

func TestSetupFlagsLoadsZoneFromEnv(t *testing.T) {
	t.Setenv("IPMAN_DNS_ZONE", "example.com")

	cmd := &Command{Self: "ipman"}

	showHelp, err := cmd.setupFlags(nil)
	if err != nil {
		t.Fatalf("setupFlags() error = %v", err)
	}
	if showHelp {
		t.Fatal("setupFlags() unexpectedly requested help")
	}

	if got := cmd.config.zone; got != "example.com" {
		t.Fatalf("setupFlags() zone = %q, want %q", got, "example.com")
	}

	if !cmd.config.v4 {
		t.Fatal("setupFlags() did not default to IPv4")
	}

	if cmd.dns == nil || cmd.ip == nil {
		t.Fatal("setupFlags() did not initialize backends")
	}
}

func TestSetupFlagsReturnsParseError(t *testing.T) {
	t.Setenv("IPMAN_DNS_ZONE", "example.com")

	cmd := &Command{Self: "ipman"}

	showHelp, err := cmd.setupFlags([]string{"-ttl=bad"})
	if err == nil {
		t.Fatal("setupFlags() error = nil, want parse failure")
	}
	if showHelp {
		t.Fatal("setupFlags() unexpectedly requested help")
	}

	if !strings.Contains(err.Error(), "parse update flags") {
		t.Fatalf("setupFlags() error = %v, want contextual parse error", err)
	}
}

func TestSetupFlagsRequiresZoneWhenUnset(t *testing.T) {
	t.Parallel()

	cmd := &Command{Self: "ipman"}

	showHelp, err := cmd.setupFlags(nil)
	if err == nil {
		t.Fatal("setupFlags() error = nil, want missing zone failure")
	}
	if showHelp {
		t.Fatal("setupFlags() unexpectedly requested help")
	}

	if err != internalerrors.ErrMissingZone {
		t.Fatalf("setupFlags() error = %v, want %v", err, internalerrors.ErrMissingZone)
	}
}
