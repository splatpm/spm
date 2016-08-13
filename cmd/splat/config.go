package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type SplatConfig struct {
	// Global configuration
	Debug bool `yaml:"debug"`
	Color bool `yaml:"color"`

	// Audit command configuration
	Audit struct {
		Report bool `yaml:"report"`
		Revert bool `yaml:"revert"`
	} `yaml:"audit"`

	// Install command configuration
	Install struct {
		Root string `yaml:"root"`
	} `yaml:"install"`

	// Remove command configuration
	Remove struct {
		Force bool `yaml:"force"`
	} `yaml:"remove"`
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
			// NOTE This will only be applicable if debug is set to true
			//      in the config file that's used.
			Debug(fmt.Sprintf("Used: %s", v))
		}
	}
	return nil
}

func ToggleBool(a bool, b *bool) {
	Debug(fmt.Sprintf("%+v %+v", a, b))
	if a == true {
		if *b == false {
			*b = true
		}
	} else {
		if *b == true {
			*b = false
		}
	}
}
