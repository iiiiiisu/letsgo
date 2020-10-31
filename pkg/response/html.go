package response

import (
	"html/template"
	"net/http"
)

func HtmlResponse(w http.ResponseWriter, tpl *template.Template, data map[string]interface{}) error {
	w.Header().Set("Content-type", "text/html")
	return tpl.Execute(w, data)
}
