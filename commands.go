package main

import (
	"os"

	"github.com/leprechau/ipman/command/check"
	"github.com/leprechau/ipman/command/update"
	"github.com/mitchellh/cli"
)

// init command factory
func initCommands() map[string]cli.CommandFactory {
	// register sub commands
	return map[string]cli.CommandFactory{
		"check": func() (cli.Command, error) {
			return &check.Command{
				Self: os.Args[0],
				Log:  log,
			}, nil
		},
		"update": func() (cli.Command, error) {
			return &update.Command{
				Self: os.Args[0],
				Log:  log,
			}, nil
		},
	}
}
