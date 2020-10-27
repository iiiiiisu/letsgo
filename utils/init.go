package utils

import (
	"fmt"
)

func init() {
	err := InitDB()
	if err != nil {
		fmt.Println(err)
	}
}
