package handler

import (
	"letsgo/utils/handler"
	"letsgo/utils/router"
	"net/http"
)

type dict = map[string]interface{}

var UrlPattern = []router.Route{
	router.Route{Path: "/",
		Handler: handler.NewRequestHandler(&IndexHandler{TplName:"templates/index.html"})},
	router.Route{Path: "/login",
		Handler: handler.NewRequestHandler(&LoginHandler{TplName:"templates/login.html"})},
	router.Route{Path: "/register",
		Handler: handler.NewRequestHandler(&RegisterHandler{TplName:"templates/register.html"})},
	router.Route{Path: "/user/(:id)",
		Handler: handler.NewRequestHandler(&UserHandler{})},
	router.Route{Path: "/static/",
		Handler: http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))},
}


func init() {
}


//func IndexHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "hello world", r.URL)
//	fmt.Fprintln(w, "Path: ", r.URL.Path)
//	fmt.Fprintln(w, "Host: ", r.URL.Host)
//	fmt.Fprintln(w, "User: ", r.URL.User.Username())
//	pwd, _ := r.URL.User.Password()
//	fmt.Fprintln(w, "Pwd: ", pwd)
//	fmt.Fprintln(w, "ForceQuery: ", r.URL.ForceQuery)
//	fmt.Fprintln(w, "Fragment: ", r.URL.Fragment)
//	fmt.Fprintln(w, "Opaque: ", r.URL.Opaque)
//	fmt.Fprintln(w, "RawPath: ", r.URL.RawPath)
//	fmt.Fprintln(w, "RawQuery: ", r.URL.RawQuery)
//	fmt.Fprintln(w, "Scheme: ", r.URL.Scheme)
//	cnt, err :=ioutil.ReadAll(r.Body)
//	if err != nil {
//		fmt.Println(err)
//	}
//	fmt.Println(string(cnt))
//	err = r.ParseForm()
//	if err != nil {
//		fmt.Println(err)
//	}
//	msg := make([]byte, 1024)
//	r.Body.Read(msg)
//	text := string(msg)
//	fmt.Println(text)
//	fmt.Fprintln(w, "Body: ", text)
//}
