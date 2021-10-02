package learn_go_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

// File Server
// Handler that used for serving static file
// https://pkg.go.dev/net/http#FileServer
func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// strip the prefix, since the url path different if not stripped
	// localhost:8080/recources/static/index.html <- 404 file not found, original
	// localhost:8080/recources/index.html <- stripped

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// use golang embed so you dont have to copy the static file
// https://pkg.go.dev/embed#FS

//go:embed resources
var resources embed.FS // embed file system

func TestFileServerEmbed(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources") // fs.Sub = open the sub directory
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
