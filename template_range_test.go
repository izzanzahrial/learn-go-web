package learn_go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

// Range loop in template
func TemplateRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Template Action If Else statement",
		"Skills": []string{
			"Focus", "Ambitious", "Lead", "Calm",
		},
	})
}

func TestTemplateRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateRange(recoder, request)

	body, err := io.ReadAll(recoder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
