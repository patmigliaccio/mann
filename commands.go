package main

import (
	"github.com/urfave/cli"
)

var commands = []cli.Command{
	{
		Name:      "add",
		Aliases:   []string{"a"},
		Usage:     "stores a new custom command",
		Action:    Add,
		ArgsUsage: `command [options]`,
	},
}
