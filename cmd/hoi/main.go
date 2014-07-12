package main

import (
	"os"
	"os/exec"

	flags "github.com/jessevdk/go-flags"
	"github.com/monochromegane/hoi"
)

var opts hoi.Options

func main() {

	args, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	if opts.Server {
		// start hoi server
		hoi.StartServer()
	} else {
		// make public
		link := hoi.MakePublic(args[0])

		// print URL
		hoi.PrintUrl(link)

		// run hoi server as a daemon
		runAsDaemon()
	}

}

func runAsDaemon() {
	cmd := exec.Command(os.Args[0], "--server")
	cmd.Start()
}
