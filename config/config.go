package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var AppConfig Config

type Config struct {
	Log   logConfig   `yaml:"log"`
	Email emailConfig `yaml:"email"`
}

type emailConfig struct {
	Host       string `yaml:"host"`
	Username   string `yaml:"username"`
	Password   string `yaml:"password"`
	PortNumber string `yaml:"portNumber"`
	Receiver   string `yaml:"receiver"`
}

type logConfig struct {
	TargetLevel string `yaml:"targetLevel"`
	Path        string `yaml:"path"`
}

func ParseFromConfig() Config {
	file, err := ioutil.ReadFile("./resources/config.yaml")
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = yaml.Unmarshal(file, &AppConfig)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return AppConfig
}
