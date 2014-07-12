package hoi

import (
	"strings"
	"testing"
)

func TestUrl(t *testing.T) {
	expect := "8081"
	server := HoiServer{
		Port: expect,
	}
	url := server.Url()
	if !strings.Contains(url, "http://") {
		t.Errorf("It should be contains %s", "http://")
	}
	if !strings.Contains(url, ":"+expect) {
		t.Errorf("It should be contains %s", ":"+expect)
	}
}
