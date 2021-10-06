package learn_go_web

import (
	"fmt"
	"net/http"
	"testing"
)

// Redirect function
// in golang so you dont have to use header manualy
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Redirections
func RedirectTo(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello Redirect")
}

// https://pkg.go.dev/net/http#Redirect
func RedirectFrom(writer http.ResponseWriter, request *http.Request) {
	http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect)
	// http.Redirect(writer, request, "/redirect-to", http.StatusPermanentRedirect)
}

func TestRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
