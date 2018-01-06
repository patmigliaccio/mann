package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/urfave/cli"
)

// Add stores a new custom command
func Add(c *cli.Context) error {
	args := c.Args()

	customCommand := ""
	for i := 0; i < len(args); i++ {
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

	return nil
}

// ParseCommandName returns the command name from cli arguments
func ParseCommandName(args cli.Args) string {
	for i := 0; i < len(args); i++ {
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
