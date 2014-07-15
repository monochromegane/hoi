package main

import (
	"fmt"
	"os"
	"os/exec"

	flags "github.com/jessevdk/go-flags"
	"github.com/monochromegane/hoi"
)

const version = "0.0.2"

var opts hoi.Options

func main() {

	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "hoi"
	parser.Usage = "[OPTIONS] PATH"
	args, err := parser.Parse()
	if err != nil {
		os.Exit(1)
	}

	hoi := hoi.NewHoi()
	switch {
	case opts.Version:
		fmt.Printf("%s\n", version)
	case opts.Clear:
		// clear all symlinks by removing public directory
		hoi.Clear()
	case opts.Server:
		// start hoi server
		hoi.Server().Start()
	default:
		if len(args) < 1 {
			parser.WriteHelp(os.Stdout)
			os.Exit(1)
		}
		// make public
		hoi.MakePublic(args[0])
		// run hoi server as a daemon
		runAsDaemon()
	}
}

func runAsDaemon() {
	cmd := exec.Command(os.Args[0], "--server")
	cmd.Start()
}
