package templates

import (
	"bufio"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

func Render(w http.ResponseWriter, data map[string]interface{}, tplName ...string) error {
	tpl, err := template.ParseFiles(tplName...)
	if err != nil {
		return err
	}
	return tpl.Execute(w, data)
}

func RenderAll(w http.ResponseWriter, data map[string]interface{}, tplName string, path string) error {
	var tpls = []string{
		tplName,
	}
	re, err := regexp.Compile(`\{\{\s*template "(.*)"\s*\}\}`)
	if err != nil {
		return err
	}
	file, err := os.Open(path + tplName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		res := re.FindStringSubmatch(line)
		if len(res) < 2 {
			continue
		}
		for _, tpl := range res[1:] {
			tpls = append(tpls, path + tpl)
		}
	}
	return Render(w, data, tpls...)
}

func Prepare(tplName string, path ...string) (*template.Template, error) {
	if len(path) == 0 {
		return template.ParseFiles(tplName)
	}
	var tpls = []string{
		tplName,
	}
	re, err := regexp.Compile(`\{\{\s*template "(.*)".*\}\}`)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(tplName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		res := re.FindStringSubmatch(line)
		if len(res) < 2 {
			continue
		}
		for _, tpl := range res[1:] {
			tpls = append(tpls, path[0] + tpl + ".html")
		}
	}
	return template.ParseFiles(tpls...)
}