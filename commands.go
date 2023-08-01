package main

import (
	stdLog "log"
	"os"

	"github.com/mitchellh/cli"

	// command implementations
	"github.com/leprechau/ipman/command/check"
	"github.com/leprechau/ipman/command/update"
)

// package global logger
var logger *stdLog.Logger

// available commands
var cliCommands map[string]cli.CommandFactory

// init command factory
func init() {
	// init logger
	logger = stdLog.New(os.Stderr, "", stdLog.LstdFlags)

	// register sub commands
	cliCommands = map[string]cli.CommandFactory{
		"check": func() (cli.Command, error) {
			return &check.Command{
				Self: os.Args[0],
				Log:  logger,
			}, nil
		},
		"update": func() (cli.Command, error) {
			return &update.Command{
				Self: os.Args[0],
				Log:  logger,
			}, nil
		},
	}
}
