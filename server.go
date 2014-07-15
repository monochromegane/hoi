package hoi

import (
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type HoiServer struct {
	DocumentRoot string
	Port         int
}

func (h HoiServer) Start() {
	http.Handle("/", http.FileServer(http.Dir(h.DocumentRoot)))
	http.ListenAndServe(":"+strconv.Itoa(h.Port), nil)
}

func (h HoiServer) Url() string {
	return "http://" + localIpAddress() + ":" + strconv.Itoa(h.Port)
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
