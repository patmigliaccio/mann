package main

import (
	"github.com/urfave/cli"
)

// Commands is the list of actionable predicates
type Commands []cli.Command

// AppCommands is the list of commands for the app
var AppCommands Commands = []cli.Command{
	{
		Name:         "add",
		Aliases:      []string{"a"},
		Usage:        "stores a new custom command",
		Action:       Add,
		ArgsUsage:    `command [options]`,
		OnUsageError: OnUsageErrorAdd,
	},
	{
		Name:      "run",
		Aliases:   []string{"r"},
		Usage:     "runs a specified command by position",
		Action:    Run,
		ArgsUsage: `command listItemPosition [options]`,
	},
}
