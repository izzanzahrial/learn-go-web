package learn_go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

// Create template function
type Person struct {
	Name string
}

func (person Person) SayHello(name string) string {
	return "Hello " + name + ", My name is " + person.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	// {{ .NameOfTheFunction "Parameter" }}
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Zahrial" }}`))
	t.ExecuteTemplate(writer, "FUNCTION", Person{
		Name: "Izzan",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateFunction(recoder, request)

	body, err := io.ReadAll(recoder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// Create global function
// https://github.com/golang/go/blob/master/src/text//template/funcs.go
func TemplateGlobalFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len .Name }}`))
	t.ExecuteTemplate(writer, "FUNCTION", Person{
		Name: "Izzan",
	})
}

func TestTemplateGlobalFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateGlobalFunction(recoder, request)

	body, err := io.ReadAll(recoder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// Create your own global function
// using Funcs
// you have to do it before parse

func TemplateCreateGlobalFunction(writer http.ResponseWriter, request *http.Request) {

	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{ // create the function
		"upper": func(value string) string { // "upper" as the name of the function
			return strings.ToUpper(value) // and anonymous function as the function
		},
	})
	t = template.Must(t.Parse(`{{upper .Name }}`)) // then parse the data
	t.ExecuteTemplate(writer, "FUNCTION", Person{
		Name: "Ahmad Izzan Zahrial",
	})
}

func TestTemplateCreateGlobalFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateCreateGlobalFunction(recoder, request)

	body, err := io.ReadAll(recoder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// Pipelines Function
// function that will run after the other
func TemplatePipelineFunction(writer http.ResponseWriter, request *http.Request) {

	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(value string) string {
			return "Hello " + value
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{ sayHello .Name | upper}}`))
	// the pipeline upper will run after sayHello is done
	// and the value from sayHello will go to the upper
	t.ExecuteTemplate(writer, "FUNCTION", Person{
		Name: "Ahmad Izzan Zahrial",
	})
}

func TestTemplatePipelineFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplatePipelineFunction(recoder, request)

	body, err := io.ReadAll(recoder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
