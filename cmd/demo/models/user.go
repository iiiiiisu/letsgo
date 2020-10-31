package models

import (
	"database/sql"
	"errors"
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

func (u *User) Logout(sId string) bool {
	if sId == "" {
		return false
	}
	rConn := redisPool.Get()
	defer rConn.Close()
	_, err := rConn.Do("set", sId, "", "EX", "1800")
	if err != nil {
		return false
	}
	return true
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

func (u *User) Get() error {
	if u.Username == "" {
		return errors.New("No key select")
	}
	rows := db.QueryRow("select nickname, avatar, gender from users where username = ?;", u.Username)
	err := rows.Scan(&u.Nickname, &u.Avatar, &u.Gender)
	return err
}
