package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"

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
	usr, _ := user.Current()
	homeDir := usr.HomeDir
	publicDir := filepath.Join(homeDir, ".hoi", "public")
	os.MkdirAll(publicDir, 0755)

	if opts.Server {
		// start hoi server
		hoi.Start(publicDir)
	} else {
		// create random directory
		random := randomString(32)
		randomDir := filepath.Join(publicDir, random)
		os.Mkdir(randomDir, 0755)

		// create symblic link
		file := args[0]
		os.Symlink(file, filepath.Join(randomDir, filepath.Base(file)))

		// print URL
		printUrl(filepath.Join(random, filepath.Base(file)))

		// run hoi server as a daemon
		cmd := exec.Command(os.Args[0], "--server")
		cmd.Start()
	}

}

func randomString(length int) string {
	alphanum := "0123456789abcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, length)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}
	return string(bytes)
}

func printUrl(path string) {
	fmt.Println(hoi.Url() + "/" + path)
}
