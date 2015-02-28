package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	flags "github.com/jessevdk/go-flags"
	"github.com/monochromegane/hoi"
)

const version = "0.1.2"

var opts hoi.Options

func main() {

	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "hoi"
	parser.Usage = "[OPTIONS] PATH|MESSAGE [@SlackUser]"
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
		args, to := parseArgs(args)
		var path string
		if abspath, patherr := hoi.TestFile(args[0]); patherr == nil {
			path = hoi.MakePublic(abspath)
		} else {
			path = hoi.MakeMessage(args)
		}
		url := hoi.ToUrl(path)
		fmt.Println(url)

		// notify
		if to != "" {
			fmt.Fprint(os.Stderr, hoi.Notify(to, url))
		}

		// run hoi server as a daemon
		runAsDaemon()
	}
}

func runAsDaemon() {
	cmd := exec.Command(os.Args[0], "--server")
	cmd.Start()
}

func parseArgs(args []string) ([]string, string) {
	if strings.HasPrefix(args[len(args)-1], "@") {
		return args[:len(args)-1], args[len(args)-1]
	}
	return args, ""
}
