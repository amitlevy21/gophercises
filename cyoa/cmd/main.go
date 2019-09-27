package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/amitlevy21/gophercises/cyoa/story"
)

func main() {
	port := flag.Int("port", 3000, "port to start cyoa web app")
	fileName := flag.String("file", "gopher.json", "JSON file containing a cyoa story")

	jsonFile, err := os.Open(*fileName)
	check(err)

	v, err := story.JSONStory(jsonFile)
	t, err := template.ParseFiles("./templates/story.html")
	if err != nil {
		log.Fatal(err)
	}

	h := story.NewHandler(v, story.WithTemplate(t), story.WithURLParseFn(customURLParser()))
	http.Handle("/", h)
	log.Printf("Listening on port %d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}

func customURLParser() func(*http.Request) string {
	return func(r *http.Request) (chapterName string) {
		path := r.URL.Path
		if path == "/" || path == "" {
			chapterName = "intro1"
		} else {
			chapterName = path[len("/story/"):]
		}
		return chapterName
	}
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}
