package server

import (
	"net"
	"net/http"
	"os"
	"regexp"
)

func Start(documentRoot string) {
	http.Handle("/", http.FileServer(http.Dir(documentRoot)))
	http.ListenAndServe(":8080", nil)
}

func Url() string {
	return "http://" + localIpAddress() + ":8080"
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
