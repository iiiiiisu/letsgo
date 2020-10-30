package parser

import (
	"errors"
	"strings"
)

type dict map[string]interface{}

func ParseStrToDict(str string) (datas dict, err error) {
	keywords := strings.Split(str, "&")
	for _, keyword := range keywords {
		data := strings.Split(keyword, "=")
		if len(data) != 2 {
			err = errors.New("Data Format Not Right. ")
			return
		}
		datas[data[0]] = data[1]
	}
	return
}
