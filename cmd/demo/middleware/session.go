package middleware

import (
	"crypto/md5"
	"fmt"
	"io"
	"letsgo/pkg"
	"letsgo/pkg/middleware"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func Md5(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

type SessionMidd struct {
	middleware.NullMiddleWare
}

func (m *SessionMidd) ProcessRequest(w http.ResponseWriter, r *http.Request) {
	// 从 request 的 cookie 中获取 sessionId 的值
	// 若不存在，则新建一个 sessionId 并写入到cookies中并返回
	// 存在则更新 sessionId 过期时间

	rConn := pkg.GetRedisPool().Get()
	defer rConn.Close()
	cookie, err := r.Cookie("sessionId")
	if err == http.ErrNoCookie {
		sessionId := m.SessionId()
		_, err := rConn.Do("set", sessionId, "", "EX", "1800")
		if err == nil {
			ck := &http.Cookie{
				Name:   "sessionId",
				Value:  sessionId,
				MaxAge: 1800,
			}
			http.SetCookie(w, ck)
		}
	} else {
		rConn.Do("EXPIRE", cookie.Value, 1800)
	}
}

func (m *SessionMidd) ProcessResponse(w http.ResponseWriter, r *http.Request) {
}

func (m *SessionMidd) SessionId() string {
	nano := time.Now().UnixNano()
	rand.Seed(nano)
	randNum := rand.Int63()
	sessionId := Md5(Md5(strconv.FormatInt(nano, 10)) + Md5(strconv.FormatInt(randNum, 10)))
	return sessionId
}
