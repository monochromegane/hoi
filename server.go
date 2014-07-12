package hoi

import (
	"net"
	"net/http"
	"os"
	"regexp"
)

type HoiServer struct {
	DocumentRoot string
	Port         string
}

func (h HoiServer) Start() {
	http.Handle("/", http.FileServer(http.Dir(h.DocumentRoot)))
	http.ListenAndServe(":"+h.Port, nil)
}

func (h HoiServer) Url() string {
	return "http://" + localIpAddress() + ":" + h.Port
}

func localIpAddress() string {
	name, _ := os.Hostname()
	addrs, _ := net.LookupHost(name)
	for _, a := range addrs {
		if m, _ := regexp.MatchString("(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})", a); m {
			return a
		}
	}
	return ""
}
