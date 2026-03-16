// Package main initializes and starts ipman
package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
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

func usageError(self, message string) string {
	return fmt.Sprintf("%s: %s\n\n%s", self, message, usage(self))
}

// main is a thin wrapper around the real application entry point.
func main() {
	os.Exit(realMain())
}

// realMain executes the application and returns a process exit code.
func realMain() int {
	logger := setupLogger()
	commands := initCommands(logger)

	if len(os.Args) < 2 {
		_, _ = os.Stderr.WriteString(usage(os.Args[0]))
		return 1
	}

	switch arg := os.Args[1]; arg {
	case "help", "-h", "--help":
		_, _ = os.Stdout.WriteString(usage(os.Args[0]))
		return 0
	default:
		if strings.HasPrefix(arg, "-") {
			message := fmt.Sprintf("unknown top-level flag %q", arg)
			if arg == "-" || arg == "--" {
				message = fmt.Sprintf("expected a subcommand before %q", arg)
			}

			_, _ = os.Stderr.WriteString(usageError(os.Args[0], message))
			return 1
		}

		if cmd, ok := commands[arg]; ok {
			return runCommand(os.Args[2:], cmd)
		}

		_, _ = os.Stderr.WriteString(usageError(os.Args[0], fmt.Sprintf("unknown command %q", arg)))
		return 1
	}
}
