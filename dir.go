package hoi

import (
	"crypto/rand"
	"os"
	"os/user"
	"path/filepath"
)

func MakePublicDir() string {
	publicDir := PublicDir()
	os.MkdirAll(publicDir, 0755)
	return publicDir
}

func PublicDir() string {
	usr, _ := user.Current()
	homeDir := usr.HomeDir
	return filepath.Join(homeDir, ".hoi", "public")
}

func LinkToFile(path string) string {

	file := filepath.Base(path)

	// create random directory
	random := randomString(32)
	randomDir := filepath.Join(PublicDir(), random)
	os.Mkdir(randomDir, 0755)

	// create symblic link
	os.Symlink(path, filepath.Join(randomDir, file))

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
