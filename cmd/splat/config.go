package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type SplatConfig struct {
	Debug bool `yaml:"debug"`
	Color bool `yaml:"color"`
}

func BuildConfig() error {
	Config = &SplatConfig{}
	for _, v := range []string{
		"/etc/spm/config.yml",
		"/etc/spm.yml",
		"./etc/spm/config.yml",
		"./etc/spm.yml",
		fmt.Sprintf("%s/.spm/config.yml", os.Getenv("HOME")),
		fmt.Sprintf("%s/.spm.yml", os.Getenv("HOME")),
	} {
		if _, err := os.Stat(v); err == nil {
			fmt.Println("Reading: ", v)
			f, e := ioutil.ReadFile(v)
			if e != nil {
				fmt.Println("ERROR: ", e.Error())
				os.Exit(1)
			}
			e = yaml.Unmarshal(f, Config)
			if e != nil {
				fmt.Println("ERROR: ", e.Error())
				os.Exit(1)
			}
		}
	}
	return nil
}
