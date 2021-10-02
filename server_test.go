package learn_go_web

import (
	"net/http" // go lang web library
	"testing"
)

// Create web server
// https://pkg.go.dev/net/http#Server
func TestServer(t *testing.T) {
	server := http.Server{ // http.server = golang web server
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe() // ListenAndServe() = run the server
	if err != nil {
		panic(err)
	}
}
