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

	app.Action = ActionHandler

	app.Run(os.Args)
}

// ActionHandler routes arguments to the appropriate methods
func ActionHandler(c *cli.Context) error {
	if c.NArg() > 0 {
		args := c.Args()

		switch action := args[0]; action {
		case "add":
			AddCommand(args)
		default:
			GetCommands(args[0])
		}

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

// AddCommand stores a new custom command to the list
func AddCommand(args cli.Args) {
	customCommand := ""
	for i := 1; i < len(args); i++ {
		customCommand += fmt.Sprintf("%s ", args[i])
	}

	filename := filepath + ParseCommandName(args) + ".yaml"
	customCommandOut := fmt.Sprintf("  - %s\r\n", customCommand)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		if err := ioutil.WriteFile(filename, []byte("cmds:\r\n"), 0644); err != nil {
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

// ParseCommandName returns the command name from cli arguments
func ParseCommandName(args cli.Args) string {
	for i := 1; i < len(args); i++ {
		arg := ParseOutCommandPrefix(args[i], "sudo")
		if arg != "" {
			return arg
		}
	}

	log.Fatal("Unable to parse arguments.")

	return ""
}

// ParseOutCommandPrefix recursively strips a prefix
// if it exists (e.g. `sudo`) and returns the first argument
func ParseOutCommandPrefix(arg string, prefix string) string {
	if strings.TrimSpace(arg) == prefix {
		return ""
	}

	prefixIndex := strings.Index(arg, prefix)
	if prefixIndex > -1 {
		return ParseOutCommandPrefix(arg[prefixIndex+len(prefix)+1:], prefix)
	}

	return strings.Fields(arg)[0]
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
