package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Auth struct {
		Login    string `yaml:login`
		Password string `yaml:password`
	} `yaml:auth`
}

var App Config

func getConf() []byte {

	yamlFile, err := ioutil.ReadFile("./src/config/config.yml")
	if err != nil {
		log.Printf("#%v ", err)
	}
	return yamlFile
}
func init() {
	err := yaml.Unmarshal(getConf(), &App)
	if err != nil {
		log.Fatalf("config error: %s", err)
	}
}
