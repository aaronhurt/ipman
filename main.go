// Package main initializes and starts ipman
package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/leprechau/ipman/command/check"
	"github.com/leprechau/ipman/command/update"
)

// setupLogger configures the application logger.
func setupLogger() *slog.Logger {
	lvl := &slog.LevelVar{} // create new level logger
	lvl.Set(slog.LevelInfo) // default to Info
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: lvl}))
	if l := os.Getenv("LOG"); l != "" {
		switch {
		case strings.HasPrefix(l, "D"):
			lvl.Set(slog.LevelDebug)
		case strings.HasPrefix(l, "W"):
			lvl.Set(slog.LevelWarn)
		case strings.HasPrefix(l, "E"):
			lvl.Set(slog.LevelError)
		}
	}

	return logger
}

// usage generates the command-line usage message.
func usage(self string) string {
	return fmt.Sprintf(`Usage: %s <command> [options]

Commands:
  check    Return current external IP address of local machine.
  update   Update DNS registry with external IP address of local machine.
  help     Show this help message.
`, self)
}

// main is a thin wrapper around the real application entry point.
func main() {
	os.Exit(realMain())
}

// realMain executes the application and returns a process exit code.
func realMain() int {
	logger := setupLogger()

	if len(os.Args) < 2 {
		_, _ = os.Stderr.WriteString(usage(os.Args[0]))
		return 1
	}

	switch os.Args[1] {
	case "check":
		cmd := &check.Command{
			Self: os.Args[0],
			Log:  logger,
		}
		return cmd.Run(os.Args[2:])
	case "update":
		cmd := &update.Command{
			Self: os.Args[0],
			Log:  logger,
		}
		return cmd.Run(os.Args[2:])
	case "help", "-h", "--help":
		_, _ = os.Stdout.WriteString(usage(os.Args[0]))
		return 0
	default:
		logger.Error("unknown command", "command", os.Args[1])
		_, _ = os.Stderr.WriteString(usage(os.Args[0]))
		return 1
	}
}
