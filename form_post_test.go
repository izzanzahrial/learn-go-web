package learn_go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Form Post
// https://pkg.go.dev/net/http#Request.ParseForm
func FormPost(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm() // in order to get the value of the post, you have to parse it first
	if err != nil {
		panic(err)
	}

	firstName := request.PostForm.Get("first_name") // Get the value
	lastName := request.PostForm.Get("last_name")

	// Actually there is a function that handle the parsing
	// request.PostFormValue("first_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=Izzan&last_name=Zahrial")
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
