package main

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type Config struct {
	BlogTitle        string
	SubTitle         string
	ShowSubtitle     bool
	Stylesheet       string
	PostLimit        int
	ShowDate         bool
	ShowAuthor       bool
	ShowSearch       bool
	SearchText       string
	SearchButtonText string
	ShowCredits      bool
	ShowPermLinks    bool
}

func LoadConfiguration() Config {
	var c Config
	b, err := ioutil.ReadFile("config/settings.toml")
	if err != nil {
		fmt.Print(err)
	}

	configfile := string(b)

	if _, err := toml.Decode(configfile, &c); err != nil {
		fmt.Println(err)
	}
	return c
}

func (c Config) GetBlogTitle() string {
	return c.BlogTitle
}

func (c Config) GetSubTitle() string {
	return c.SubTitle
}

func (c Config) GetShowSubTitle() bool {
	return c.ShowSubtitle
}

func (c Config) GetStylesheet() string {
	return c.Stylesheet
}

func (c Config) GetPostLimit() int {
	return c.PostLimit
}

func (c Config) GetShowDate() bool {
	return c.ShowDate
}

func (c Config) GetShowAuthor() bool {
	return c.ShowAuthor
}

func (c Config) GetShowSearch() bool {
	return c.ShowSearch
}

func (c Config) GetSearchText() string {
	return c.SearchText
}

func (c Config) GetSearchButtonText() string {
	return c.SearchButtonText
}

func (c Config) GetShowCredits() bool {
	return c.ShowCredits
}

func (c Config) GetShowPermLinks() bool {
	return c.ShowPermLinks
}
