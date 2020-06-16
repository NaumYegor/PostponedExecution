package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var (
	confPath string = "config.yaml"
)

type YamlConf struct {
	Config struct {
		Port uint16 `yaml:"port"`
	}
}

func (conf *YamlConf) getConf() *YamlConf {
	yamlFile, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Printf("config-yaml-file error \n%e", err)
	}

	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		log.Printf("error in unmarshaling config\n%e", err)
	}

	return conf
}

func main() {
	var conf YamlConf
	conf.getConf()

	fmt.Printf("The port is %d", conf.Config.Port)
}
