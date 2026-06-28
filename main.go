package main

import (
	"fmt"
	"label-ping/cli"
	"os"
)

func main() {
	cfg, err := cli.ParseArgs()
	if err != nil {
		cli.PrintHeader()
		fmt.Printf("[\x1b[31mError\x1b[0m]: %v\n\n", err)
		cli.PrintUsage()
		os.Exit(1)
	}

	os.Exit(cli.Run(cfg))
}
