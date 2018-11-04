package main

import (
	"fmt"
)

func main() {
	p := BlogPost{id: 4124, title: "Testeintrag", author: "Marc", content: "whdawoudhawud"}

	fmt.Println(p.GetID())

}
