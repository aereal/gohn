package main

import (
	"flag"
	"fmt"
	"io"
	"os"
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

	if result, err := Parse(os.Stdin); err != nil {
		fmt.Printf("! Error: %#v\n", err)
		return 1
	} else {
		fmt.Printf("%#v\n", result)
		return 0
	}
}
