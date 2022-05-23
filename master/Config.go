package master

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ApiPort         int `json:"apiPort"`
	ApiReadTimeout  int `json:"apiReadTimeout"`
	ApiWriteTimeout int `json:"apiWriteTimeout"`
}

var (
	G_config *Config
)

func InitConfig(filename string) (err error) {

	var (
		content []byte
		conf    Config
	)

	content, err = ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(content, &conf)
	if err != nil {
		return err
	}

	G_config = &conf
	return
}
