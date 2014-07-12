package hoi

import (
	"crypto/rand"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func MakePublic(file string) string {
	// create hoi public directory
	publicDir := MakePublicDir()

	// create symblic link
	return linkToFile(file, publicDir)
}

func StartServer() {
	Start(publicDir())
}

func PrintUrl(path string) {
	fmt.Println(Url() + "/" + path)
}

func MakePublicDir() string {
	publicDir := publicDir()
	os.MkdirAll(publicDir, 0755)
	return publicDir
}

func publicDir() string {
	usr, _ := user.Current()
	homeDir := usr.HomeDir
	return filepath.Join(homeDir, ".hoi", "public")
}

func linkToFile(src, dest string) string {

	file := filepath.Base(src)

	// create random directory
	random := randomString(32)
	randomDir := filepath.Join(dest, random)
	os.Mkdir(randomDir, 0755)

	// create symblic link
	os.Symlink(src, filepath.Join(randomDir, file))

	return filepath.Join(random, file)
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
