package main

import (
	"github.com/urfave/cli"
)

var commands = []cli.Command{
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
