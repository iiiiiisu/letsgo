package handler

import (
	"html/template"
	"net/http"
)

type TemplateHandler struct {
	BaseHandler
	TplName string
}

func (h *TemplateHandler) Render(w http.ResponseWriter, data map[string]interface{}) error {
	tpl, err := template.ParseFiles(h.TplName)
	if err != nil {
		return err
	}
	err = tpl.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}