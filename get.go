package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kylelemons/go-gypsy/yaml"
)

// Get retrieves and prints the list of commands
func Get(command string) {
	cmds, err := GetCommands(command)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("\r\n	Name: %s \r\n\r\n", command)

	cmdOut := "	Commands:"
	for i, cmd := range cmds {
		cmdOut += "\r\n" + fmt.Sprintf(`	%v.	%s`, i+1, cmd)
	}

	fmt.Println(cmdOut)
	fmt.Println("")
}

// GetCommands retrieves the list of commands
func GetCommands(command string) ([]string, error) {
	var cmds []string

	filename := filepath + command + ".yaml"

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, fmt.Errorf("%q has no commands listed. (%q)", command, filename)
	}

	config, err := yaml.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("%q: %s", filename, err)
	}

	node, err := yaml.Child(config.Root, "cmds")
	if err != nil {
		return nil, fmt.Errorf("%q: %s", filename, err)
	}

	for _, cmd := range node.(yaml.List) {
		cmdString := fmt.Sprintf("%s", cmd.(yaml.Scalar))
		cmds = append(cmds, cmdString)
	}

	return cmds, nil
}

// GetCustomCommand returns the command by list position
func GetCustomCommand(cmd string, k int) string {
	cmds, err := GetCommands(cmd)
	if err != nil {
		log.Fatal(err)
	}

	if len(cmds) < k || k < 1 {
		log.Fatalf("%v is not a valid position.", k)
	}

	return cmds[k-1]
}
