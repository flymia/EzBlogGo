package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateID(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func main() {
	c := LoadConfiguration()
	fmt.Println(c.GetBlogTitle())

	fmt.Println("------")

	UserTest()
}
