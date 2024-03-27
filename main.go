// Package main initializes and starts ipman
package main

import (
	"fmt"
	"os"

	// CLI library
	"github.com/mitchellh/cli"
)

// it all starts here
func main() {
	var c *cli.CLI // cli object
	var status int // exit status
	var err error  // general error holder

	// init and populate cli object
	c = cli.NewCLI(appName, appVersion)
	c.Args = os.Args[1:]       // arguments minus command
	c.Commands = initComands() // see commands.go

	// run command and check return
	if status, err = c.Run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error executing CLI: %s\n", err)
	}

	// exit
	os.Exit(status)
}
