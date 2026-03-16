// Package main initializes and starts ipman
package main

import (
	"log/slog"
	"os"
	"strings"

	// CLI library
	"github.com/mitchellh/cli"
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

// it all starts here
func main() {
	var c *cli.CLI // cli object
	var status int // exit status
	var err error  // general error holder
	logger := setupLogger()

	// init and populate cli object
	c = cli.NewCLI(appName, appVersion)
	c.Args = os.Args[1:]              // arguments minus command
	c.Commands = initCommands(logger) // see commands.go

	// run command and check return
	if status, err = c.Run(); err != nil {
		logger.Error("error executing CLI", "err", err)
	}

	os.Exit(status)
}
