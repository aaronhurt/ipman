package main

import (
	"log/slog"
	"os"

	"github.com/leprechau/ipman/command/check"
	"github.com/leprechau/ipman/command/update"
	"github.com/mitchellh/cli"
)

// init command factory
func initCommands(logger *slog.Logger) map[string]cli.CommandFactory {
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
