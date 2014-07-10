package main

import (
	"fmt"
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

	// create hoi public directory
	publicDir := hoi.MakePublicDir()

	if opts.Server {
		// start hoi server
		hoi.Start(publicDir)
	} else {
		// create symblic link
		link := hoi.LinkToFile(args[0])

		// print URL
		printUrl(link)

		// run hoi server as a daemon
		cmd := exec.Command(os.Args[0], "--server")
		cmd.Start()
	}

}

func printUrl(path string) {
	fmt.Println(hoi.Url() + "/" + path)
}
