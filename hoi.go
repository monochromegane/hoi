package hoi

import (
	"crypto/rand"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

type Hoi struct {
	publicDir string
	config    Config
}

func NewHoi() *Hoi {
	// create public dir
	publicDir := publicDir()
	os.MkdirAll(publicDir, 0755)

	return &Hoi{
		publicDir: publicDir,
		config:    Load(configPath()),
	}
}

func (h Hoi) MakePublic(file string) string {
	linked := h.makePublic(file, h.publicDir)
	h.printUrl(linked)
	return linked
}

func (h Hoi) makePublic(src, dest string) string {
	// create random directory
	random := randomString(32)
	randomDir := filepath.Join(dest, random)
	os.Mkdir(randomDir, 0755)

	// make public by symblic link
	file := filepath.Base(src)
	os.Symlink(src, filepath.Join(randomDir, file))

	return filepath.Join(random, file)
}

func (h Hoi) Server() *HoiServer {
	return &HoiServer{
		DocumentRoot: h.publicDir,
		Port:         h.config.Port,
	}
}

func (h Hoi) Clear() {
	os.RemoveAll(h.publicDir)
}

func (h Hoi) printUrl(path string) {
	server := h.Server()
	fmt.Println(server.Url() + "/" + path)
}

func publicDir() string {
	return filepath.Join(homeDir(), ".hoi", "public")
}

func configPath() string {
	return filepath.Join(homeDir(), ".hoi", "conf.json")
}

func homeDir() string {
	usr, _ := user.Current()
	return usr.HomeDir
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
