package main

import (
	"math/rand"
	"time"
)

type BlogPost struct {
	id      int
	title   string
	author  string
	content string
}

func (b BlogPost) GenerateID(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

//func (b BlogPost) ExportHTML(post []byte) []byte {
//
//}

func (b BlogPost) GetID() int {
	return b.id
}

func (b BlogPost) GetTitle() string {
	return b.title
}

func (b BlogPost) GetAuthor() string {
	return b.author
}
