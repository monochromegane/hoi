package hoi

import (
	"net"
	"net/http"
	"regexp"
	"strings"
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
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		a := addr.String()
		if m, _ := regexp.MatchString("(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})\\.(\\d{1,3})", a); m {
			if strings.Contains(a, "127.0.0.1") {
				continue
			}
			return strings.Split(a, "/")[0]
		}
	}
	return ""
}
