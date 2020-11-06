package pkg

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(driverName string, user string, pwd string, host string, port string, name string) error {
	var err error
	dataSource := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8`, user, pwd, host, port, name)
	db, err = sql.Open(driverName, dataSource)
	if err != nil {
		return err
	}
	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	return nil
}

func GetDB() *sql.DB {
	return db
}

func CloseDB() error {
	var err error
	if db != nil {
		err = db.Close()
		return err
	}
	return nil
}
