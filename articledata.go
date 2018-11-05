package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// our struct which contains the complete
// array of all Articles in the file
type Articles struct {
	XMLName  xml.Name  `xml:"articles"`
	Articles []Article `xml:"article"`
}

// the Article struct, this contains our
// Type attribute, our Article's name and
// a social struct which will contain all
// our social links
type Article struct {
	XMLName xml.Name `xml:"article"`
	ID      string   `xml:"id,attr"`
	Title   string   `xml:"title"`
	Author  string   `xml:"author"`
	Date    string   `xml:"date"`
	Content string   `xml:"content"`
}

func UserTest() {

	// Open our xmlFile
	xmlFile, err := os.Open("articles/articledata.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened Articles.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Articles array
	var Articles Articles
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'Articles' which we defined above
	xml.Unmarshal(byteValue, &Articles)

	fmt.Println()

	// we iterate through every user within our Articles array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(Articles.Articles); i++ {
		fmt.Println("Article ID: " + Articles.Articles[i].ID)
		fmt.Println("Titel: " + Articles.Articles[i].Title)
		fmt.Println("Autor: " + Articles.Articles[i].Author)
		fmt.Println("Date: " + Articles.Articles[i].Date)
		fmt.Println("Content: " + Articles.Articles[i].Content)

		fmt.Println()
	}

}
