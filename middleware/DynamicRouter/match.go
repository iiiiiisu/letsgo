package DynamicRouter

import (
	"letsgo/handler"
	"letsgo/utils/middleware"
	"net/http"
	"regexp"
	"strings"
)

type DynamicRouterParserMidd struct {
	middleware.NullMiddleWare
}

func (m *DynamicRouterParserMidd) ProcessRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	for _, route := range handler.UrlPattern {
		isMatch, datas := pathMatch(path, route.Path)
		if isMatch && datas != "" {
			r.URL.Path = route.Path
			r.Header.Set("UrlValues", datas)
			return
		}
	}
}

func pathMatch(path string, route string) (isMatch bool, datas string) {
	pathList := strings.Split(path, "/")
	routeList := strings.Split(route, "/")
	length := len(pathList)
	if length != len(routeList) {
		isMatch = false
		return
	}
	for i:= 1; i < length; i++ {
		getMatched, data :=match(pathList[i], routeList[i])
		if !getMatched {
			isMatch = false
			return
		}  else if data != "" {
			if datas == "" {
				datas += "&"
			}
			datas += data
		}
	}
	return true, datas
}


func match(a string, b string) (isMatch bool, data string) {
	if ok, _ := regexp.MatchString(`^[a-zA-Z0-9]*\(:[a-zA-Z]*\)[a-zA-Z0-9]*$`, b); !ok {
		if a == b {
			isMatch = true
		} else {
			isMatch = false
		}
		return
	}
	reg := regexp.MustCompile(`\(:[a-zA-Z]*\)`)
	keyRegStr := reg.ReplaceAllString(b, `\(:(.*)\)`)
	keyReg := regexp.MustCompile(keyRegStr)
	key := keyReg.FindStringSubmatch(b)
	if len(key) != 2 {
		isMatch = false
		return
	}
	valueRegStr := reg.ReplaceAllString(b, `([a-zA-Z0-9]*)`)
	valueReg := regexp.MustCompile(`^` + valueRegStr + `$`)
	value := valueReg.FindStringSubmatch(a)
	if len(value) != 2 {
		isMatch = false
		return
	}
	data = key[1] + "=" + value[1]
	isMatch = true
	return
}