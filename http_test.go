package learn_go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprint(writer, "Hello world")
}

// Http Test = test the http without running the web server
// https://pkg.go.dev/net/http/httptest
func TestHttp(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil) // create test request
	recoder := httptest.NewRecorder()                                                  // to record the result

	HelloHandler(recoder, request)

	response := recoder.Result() // the result
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	bodyString := string(body)

	fmt.Println(bodyString)
}
