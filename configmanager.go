package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"os/ioutil"
)

type Config struct {
	title      string
	stylesheet string
	postlimit  int
	dateshow   bool
}

func Test() {
	b, err := ioutil.ReadFile("config/settings.toml")
	if err != nil {
		fmt.Print(err)
	}

	configfile := string(b)

	var conf Config
	if _, err := toml.Decode(configfile, &conf); err != nil {
		fmt.Println(err)
	}

}
