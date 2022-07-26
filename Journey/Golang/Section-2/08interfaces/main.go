package main

import (
	"fmt"
	"log"
)

// Defining an interface:

type Stringer interface {
	String() string
}

// Defining a type:

type Article struct {
	Title  string
	Author string
}

// Creating a receiver function:

func (a Article) String() string {
	return fmt.Sprintf("The %q article was written by %s.", a.Title, a.Author)
}

func main() {

	// creating an instance of 'Article' type => "anArticle"
	anArticle := Article{
		Title:  "Interfaces in Go",
		Author: "Digital Ocean",
	}

	log.Println(anArticle.String())

}
