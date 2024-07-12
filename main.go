// Package main initializes and starts ipman
package main

import (
	"fmt"
	"log/slog"
	"os"
	"strings"

	// CLI library
	"github.com/mitchellh/cli"
)

// log is a package global logger
var log *slog.Logger

// setupLogger configures the package global logger
func setupLogger() {
	lvl := &slog.LevelVar{} // create new level logger
	lvl.Set(slog.LevelInfo) // default to Info
	log = slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: lvl}))
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
}

// it all starts here
func main() {
	var c *cli.CLI // cli object
	var status int // exit status
	var err error  // general error holder

	// setup logger
	setupLogger()

	// init and populate cli object
	c = cli.NewCLI(appName, appVersion)
	c.Args = os.Args[1:]        // arguments minus command
	c.Commands = initCommands() // see commands.go

	// run command and check return
	if status, err = c.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err)
	}

	// exit
	os.Exit(status)
}
