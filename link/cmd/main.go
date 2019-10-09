package main

import (
	"fmt"
	"github.com/amitlevy21/gophercises/link/parser"
	"os"
)

func main() {
	r, err := os.Open("ex1.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	parser.Parse(r)
}
