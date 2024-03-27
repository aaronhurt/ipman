package main

import (
	stdLog "log"
	"os"

	"github.com/mitchellh/cli"

	// command implementations
	"github.com/leprechau/ipman/command/check"
	"github.com/leprechau/ipman/command/update"
)

// command logger
var logger = stdLog.New(os.Stderr, "", stdLog.LstdFlags)

// init command factory
func initComands() map[string]cli.CommandFactory {
	// register sub commands
	return map[string]cli.CommandFactory{
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
