package learn_go_web

import (
	"fmt"
	"net/http"
	"testing"
)

// Middleware
// a way so you dont have to create or use the same handler at the start every handler
// like check login, catch error

type LogMiddleware struct { // using http.Handler interface
	Handler http.Handler
}

func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Before Execute Handler")
	middleware.Handler.ServeHTTP(writer, request)
	fmt.Println("After Execute Handler")
}

type ErrorHandler struct {
	Handler http.Handler
}

func (errorHandler *ErrorHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	defer func() { // create defer func to catch the error
		err := recover()
		if err != nil {
			fmt.Println("Error")
			writer.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(writer, "Error : %s", err)
		}
	}()
	errorHandler.Handler.ServeHTTP(writer, request)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Handler Executed")
		fmt.Fprint(writer, "Hello Middleware")
	})
	mux.HandleFunc("/panic", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("Error Executed")
		panic("Panic")
	})

	logMiddleware := &LogMiddleware{
		Handler: mux,
	}

	ErrorHandler := &ErrorHandler{
		Handler: logMiddleware,
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: ErrorHandler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
