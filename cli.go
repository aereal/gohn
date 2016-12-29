package main

import (
	"flag"
	"fmt"
	"io"
)

const Name string = "text-hatena"

type CLI struct {
	outStream, errorStream io.Writer
}

func (cli *CLI) Run(args []string) int {
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errorStream)

	if err := flags.Parse(args[1:]); err != nil {
		return 1
	}

	parsedArgs := flags.Args()
	if len(parsedArgs) < 1 {
		fmt.Fprintln(cli.errorStream, "Usage: "+Name)
		return 1
	}

	return 0
}
