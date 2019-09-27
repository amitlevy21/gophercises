package story

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type handler struct {
	s         Story
	t         *template.Template
	URLParser func(*http.Request) string
}

// HandlerOption enables the use of functional options
// when creating a handler, you may pass a Handler option for
// configuring a custom template, see WithTemplate function
type HandlerOption func(*handler)

// NewHandler serves the story and its options
func NewHandler(s Story, opts ...HandlerOption) http.Handler {
	h := defaultHandler(s)
	applyHandlerOptions(&h, opts...)
	return h
}

func defaultHandler(s Story) handler {
	t, err := template.ParseFiles("./templates/story.html")
	if err != nil {
		log.Fatal(err)
	}
	p := defaultURLParser()
	return handler{s, t, p}
}

func defaultURLParser() func(*http.Request) string {
	return func(r *http.Request) (chapterName string) {
		path := r.URL.Path
		if path == "/" || path == "" {
			chapterName = "intro"
		} else {
			chapterName = path[len("/story/"):]
		}
		return chapterName
	}
}

func applyHandlerOptions(h *handler, opts ...HandlerOption) {
	// overwrite defaults with given options
	for _, opt := range opts {
		opt(h)
	}
}

// WithTemplate is a functional option for configuring a custom template
func WithTemplate(t *template.Template) func(h *handler) {
	return func(h *handler) {
		h.t = t
	}
}

// WithURLParseFn is a functional option for custom parsing of the received URL
// to the story thats need to be shows to the user
func WithURLParseFn(parseFn func(*http.Request) string) func(h *handler) {
	return func(h *handler) {
		h.URLParser = parseFn
	}
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	chapterName := h.URLParser(r)
	h.feedChapter(w, chapterName)
}

func (h handler) feedChapter(w http.ResponseWriter, chapterName string) {
	if chap, ok := h.s[chapterName]; ok {
		log.Printf("Feeding chapter %s", chapterName)
		err := h.t.Execute(w, chap)
		if err != nil {
			http.Error(w, "Something went wrong", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, fmt.Sprintf("Chapter %s not found.", chapterName), http.StatusNotFound)
	}
}
