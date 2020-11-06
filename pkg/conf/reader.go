package conf

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Read() (config Config, err error) {
	var (
		path string
		f    *os.File
		text []byte
	)
	if path, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		return
	}
	if f, err = os.Open(path + `\conf\config.json`); err != nil {
		return
	}
	if text, err = ioutil.ReadAll(f); err != nil {
		return
	}
	if err = json.Unmarshal(text, &config); err != nil {
		return
	}
	return
}
