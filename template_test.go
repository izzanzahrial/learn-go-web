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

// Template
// a dynamic way to serve data in golang html
// you can either make it from string or file
// https://pkg.go.dev/html/template
func TemplateHTML(writer http.ResponseWriter, request *http.Request) {
	templateText := "<html><body>{{.}}</body></html>" // create the string
	// the dynamic data gonna go the "{{.}}"
	t, err := template.New("SIMPLE").Parse(templateText)
	if err != nil {
		panic(err)
	}

	// t := template.Must("SIMPLE").Parse(templateText) <- template.Must with build in err

	t.ExecuteTemplate(writer, "SIMPLE", "Hello HTML Template")

}

func TestTemplateHTML(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateHTML(recoder, request)

	body, err := io.ReadAll(recoder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// Template using file
// use .gohtml for file type
// https://pkg.go.dev/html/template#Template.ParseFiles
func TemplateHTMLFile(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFiles("./templates/index.gohtml")
	if err != nil {
		panic(err)
	}
	// t := template.Must(template.ParseFiles("./templates/index.gohtml"))

	t.ExecuteTemplate(writer, "index.gohtml", "Hello HTML Template File")
}

func TestTemplateHTMLFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateHTMLFile(recoder, request)

	body, err := io.ReadAll(recoder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// load all .gohtml in the directory using ParseGlob
func TemplateHTMLDirectory(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseGlob("./templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	// t := template.Must(template.ParseFiles("./templates/index.gohtml"))

	t.ExecuteTemplate(writer, "index.gohtml", "Hello HTML Template Directory")
}

func TestTemplateHTMLDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateHTMLDirectory(recoder, request)

	body, err := io.ReadAll(recoder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// using go embed

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateHTMLEmbed(writer http.ResponseWriter, request *http.Request) {
	t, err := template.ParseFS(templates, "templates/*.gohtml")
	if err != nil {
		panic(err)
	}
	// t := template.Must(template.ParseFiles("./templates/index.gohtml"))

	t.ExecuteTemplate(writer, "index.gohtml", "Hello HTML Template Embed")
}

func TestTemplateHTMLEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateHTMLEmbed(recoder, request)

	body, err := io.ReadAll(recoder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
