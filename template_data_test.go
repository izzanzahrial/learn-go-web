package learn_go_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Create Template using Map
func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Izzan",
		"Address": map[string]interface{}{
			"Street": "Test street Map",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateDataMap(recoder, request)

	body, err := io.ReadAll(recoder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

// Create Template using Struct
type Page struct {
	Title   string
	Name    string
	Address Address
}

type Address struct {
	Street string
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Izzan",
		Address: Address{
			Street: "Test street Struct",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateDataStruct(recoder, request)

	body, err := io.ReadAll(recoder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
