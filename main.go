package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"strings"

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

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		os.MkdirAll(filepath, 0777)
	}

	app.Action = func(c *cli.Context) error {
		if c.NArg() > 0 {
			args := c.Args()
			if args[0] != "add" {
				GetCommands(args[0])
			} else {
				AddCommand(args)
			}

		}

		return nil
	}

	app.Run(os.Args)
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

// AddCommand stores a new custom command to the list
func AddCommand(args cli.Args) {
	customCommand := ""
	for i := 1; i < len(args); i++ {
		customCommand += fmt.Sprintf("%s ", args[i])
	}

	filename := filepath + strings.Fields(args[1])[0] + ".yaml"
	customCommandOut := fmt.Sprintf("\r\n  - %s", customCommand)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := ioutil.WriteFile(filename, []byte("cmds:"), 0644); err != nil {
			log.Fatal(err)
		}
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	if _, err = f.WriteString(customCommandOut); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Added: " + customCommand)
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
