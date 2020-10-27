package handler

import (
	"database/sql"
	"fmt"
	"letsgo/models"
	"letsgo/utils"
	"letsgo/utils/cache"
	"letsgo/utils/handler"
	"letsgo/utils/templates"
	"net/http"
)

type LoginHandler struct {
	handler.BaseHandler
	TplName string
}


func (h *LoginHandler) Get(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"username": "",
		"password": "",
	}
	err := templates.Render(w, h.TplName, data)
	if err != nil {
		fmt.Println(err)
	}
}


func (h *LoginHandler) Post(w http.ResponseWriter, r *http.Request) {r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	data := map[string]interface{}{
		"username": username,
		"password": password,
	}
	if username == "" || password == "" {
		data["err_msg"] = "请输入账号和密码"
	} else {
		db := utils.GetDB()
		rows := db.QueryRow("select username, pwd from users where username = ?;", username)
		u := models.User{}
		err := rows.Scan(&u.Username, &u.Password)
		if err == sql.ErrNoRows || u.Password != password {
			data["err_msg"] = "账号或密码错误"
		} else {
			cookie, _ := r.Cookie("sessionId")
			if cookie != nil {
				rConn := cache.Pool.Get()
				defer rConn.Close()
				rConn.Do("set", cookie.Value, username)
			}
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
	}
	err := templates.Render(w, h.TplName, data)
	if err != nil {
		fmt.Println(err)
	}
	return
}
