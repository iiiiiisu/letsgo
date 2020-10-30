package models

import (
	"database/sql"
)

type User struct {
	Id       int
	Username string
	Nickname string
	Password string
	Avatar   string
	Gender   byte
}

func (u *User) Login(username string, pwd string) bool {
	rows := db.QueryRow("select username, pwd from users where username = ?;", username)
	err := rows.Scan(&u.Username, &u.Password)
	if err == sql.ErrNoRows || u.Password != pwd {
		return false
	}
	return true
}

func (u *User) Logout() {

}

func (u *User) Register(username string, pwd string, nickname string) bool {
	res, err := db.Exec(`INSERT INTO users (username, pwd, nickname) VALUES (?, ?, ?)`,
		username, pwd, nickname)
	if err != nil {
		return false
	}
	if _, err := res.LastInsertId(); err != nil {
		return false
	}
	u.Username = username
	u.Password = pwd
	u.Nickname = nickname
	return true
}

func (u *User) SetSession(sId string) bool {
	if sId == "" || u.Username == "" {
		return false
	}
	rConn := redisPool.Get()
	defer rConn.Close()
	rConn.Do("set", sId, u.Username)
	return true
}
