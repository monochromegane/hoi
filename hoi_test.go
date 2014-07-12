package hoi

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMakePublic(t *testing.T) {
	publicDir := "public_test"
	defer os.RemoveAll(publicDir)

	hoi := Hoi{}
	linked := hoi.makePublic("hoi.go", publicDir)

	file, err := os.Lstat(filepath.Join(publicDir, linked))
	if err != nil {
		t.Errorf("It should be made symlink %s, %s", file, err)
	}
}

func TestRandomString(t *testing.T) {
	expect := 10
	random := randomString(expect)
	if len(random) != 10 {
		t.Errorf("It should equal %d", expect)
	}
}
