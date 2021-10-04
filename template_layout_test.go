package learn_go_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

// create layout template
// in order to use layout, you have to include all the template that you use or
// Parse blob or embed all the files
// because if you use parse files, you have to input it one by one
func TemplateLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/layout.gohtml",
		"./templates/footer.gohtml",
	))
	// i can use layout not layout.gohtml because using define, check layout.gohtml
	t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "Template Layout Header and Footer",
		"Name":  "Izzan",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recoder := httptest.NewRecorder()

	TemplateLayout(recoder, request)

	body, err := io.ReadAll(recoder.Result().Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}
