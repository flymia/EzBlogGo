package main

// TODO : Add date

import (
	"gopkg.in/russross/blackfriday.v2"
	"math/rand"
	"time"
)


type BlogPost struct {
	id      int
	date    string
	title   string
	author  string
	content string
}

func (b BlogPost) GenerateID(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func (b BlogPost) ExportHTML() string {
	//text := []byte(b.content)
	//converted := github_flavored_markdown.Markdown(text)
	//converteds := string(converted[:])
	//return converteds

	text := []byte(b.content)
	converted := blackfriday.Run(text)
	converteds := string(converted[:])
	return converteds
}

func (b BlogPost) GetID() int {
	return b.id
}

func (b BlogPost) GetTitle() string {
	return b.title
}

func (b BlogPost) GetAuthor() string {
	return b.author
}

func (b BlogPost) GetContent() string {
	return b.content
}

func (b BlogPost) GetDate() string {
	return b.date
}
