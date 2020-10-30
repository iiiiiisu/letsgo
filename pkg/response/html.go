package response

import (
	"letsgo/pkg/templates"
	"net/http"
)

func HtmlResponse(w http.ResponseWriter, tplName string, data map[string]interface{}) {
	w.Header().Set("Content-type", "text/html")
	templates.Render(w, tplName, data)
}
