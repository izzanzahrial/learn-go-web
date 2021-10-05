package learn_go_web

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Template caching

//go:embed templates/*.gohtml
var templates2 embed.FS

// is better to make the template out of the function as a global variable
// by doing this, we only call the template and parse once
// save in memory
var myTemplates = template.Must(template.ParseFS(templates2, "templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request) {
	myTemplates.ExecuteTemplate(writer, "index.gohtml", "Hello HTML Template")
}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateCaching(recoder, request)

	body, err := io.ReadAll(recoder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
