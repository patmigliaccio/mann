package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

// CheckArgsLen throws errors if min and max arguments conditions are not met
func CheckArgsLen(c *cli.Context, min int, max int) {
	argCount := c.NArg()
	if max < min {
		log.Fatalf("%s: arguments (max < min) \r\n", c.Command.Name)
	}

	if argCount < min || argCount > max {
		if argCount > 1 {
			fmt.Printf("%v too many arguments specified.", argCount-max)
		} else {
			fmt.Printf("%v too little arguments specified.", min-argCount)
		}

		fmt.Printf(" Use 'mann --help' for more information.\r\n")

		os.Exit(1)
	}
}
