package main

import "os"

func main() {
	cli := &CLI{outStream: os.Stdout, errorStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}
