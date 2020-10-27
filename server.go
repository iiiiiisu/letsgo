package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	defer StopServe()
	ListenAndServe("localhost:8000")
}


