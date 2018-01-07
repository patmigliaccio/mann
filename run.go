package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/urfave/cli"
)

// Run executes the specified command in the list
func Run(c *cli.Context) error {
	args := c.Args()

	CheckArgsLen(c, 2, 2)

	k, err := strconv.ParseInt(args[1], 10, 0)
	if err != nil {
		log.Fatalf("%v is not a valid integer.", k)
	}

	cmd, err := GetCustomCommand(args[0], int(k))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(cmd)

	err = ExecCommand(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

// ExecCommand runs the specified command in the shell
func ExecCommand(cmdString string) error {
	cmd := exec.Command("sh", "-c", cmdString)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
