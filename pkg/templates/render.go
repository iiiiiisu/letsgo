package templates

import (
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, tplName string, data map[string]interface{}) error {
	tpl, err := template.ParseFiles(tplName)
	if err != nil {
		return err
	}
	err = tpl.Execute(w, data)
	return err
}
