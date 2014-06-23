package server

import (
	"net/http"
)

func Start(documentRoot string) {
	http.Handle("/", http.FileServer(http.Dir(documentRoot)))
	http.ListenAndServe(":8080", nil)
}
