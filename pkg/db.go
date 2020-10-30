package pkg

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/gd?charset=utf8")
	if err != nil {
		panic(err)
	}
	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() error {
	err := db.Close()
	return err
}
