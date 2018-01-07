package main

import (
	"os"
	"runtime"

	"github.com/urfave/cli"
)

var (
	filepath = UserHomeDir() + `/.mann/`
)

func main() {
	app := cli.NewApp()
	app.Name = "mann"
	app.Usage = "your personal man pages"
	app.UsageText = "mann [command] os-command [command options]"
	app.Version = "0.3.0"
	app.Authors = []cli.Author{
		{
			Name:  "Pat Migliaccio",
			Email: "pat@patmigliaccio.com",
		},
	}

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		os.MkdirAll(filepath, 0777)
	}

	app.Commands = commands
	app.Action = ActionHandler

	app.Run(os.Args)
}

// ActionHandler routes arguments to the appropriate methods
func ActionHandler(c *cli.Context) error {
	if c.NArg() > 0 {
		args := c.Args()
		Get(args[0])
	}

	return nil
}

// UserHomeDir returns the home directory of the user cross platform
func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
