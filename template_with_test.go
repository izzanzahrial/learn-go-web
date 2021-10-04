package learn_go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

// access the nested value in data using with template
func TemplateWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))
	t.ExecuteTemplate(writer, "with.gohtml", map[string]interface{}{
		"Title": "Template Action If Else statement",
		"Bio": map[string]interface{}{
			"Name":        "Izzan",
			"Age":         26,
			"Nationality": "Indonesia",
		},
	})
}

func TestTemplateWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateWith(recoder, request)

	body, err := io.ReadAll(recoder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
