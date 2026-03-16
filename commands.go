package main

import (
	"log/slog"
	"os"

	"github.com/leprechau/ipman/command/check"
	"github.com/leprechau/ipman/command/update"
)

type runner interface {
	Run([]string) int
}

type commandFactory func() runner

func initCommands(logger *slog.Logger) map[string]commandFactory {
	return map[string]commandFactory{
		"check": func() runner {
			return &check.Command{
				Self: os.Args[0],
				Log:  logger,
			}
		},
		"update": func() runner {
			return &update.Command{
				Self: os.Args[0],
				Log:  logger,
			}
		},
	}
}

func runCommand(args []string, newCommand commandFactory) int {
	return newCommand().Run(args)
}
