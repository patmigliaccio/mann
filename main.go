package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/kylelemons/go-gypsy/yaml"
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
	app.Version = "0.2.0"
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
		GetCommands(args[0])
	}

	return nil
}

// GetCommands retrieves the list of commands
func GetCommands(command string) {
	filename := filepath + command + ".yaml"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		fmt.Printf("%q has no commands listed. \r\n", command)
		os.Exit(1)
	}

	config, err := yaml.ReadFile(filename)
	if err != nil {
		fmt.Printf("%q: %s", filename, err)
	}

	fmt.Printf("\r\n	Name: %s \r\n\r\n", command)

	node, _ := yaml.Child(config.Root, "cmds")

	cmdOut := "	Commands: \r\n"
	for _, cmd := range node.(yaml.List) {
		cmdOut += fmt.Sprintf(`		%s 
`, cmd.(yaml.Scalar))
	}

	fmt.Println(cmdOut)
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
