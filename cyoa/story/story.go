package story

import (
	"encoding/json"
	"io"
)

// Story maps from chapter name to a chapter struct
type Story map[string]chapter

type chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []option `json:"options"`
}

type option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// JSONStory reads from the given fileName a story in a JSON format
func JSONStory(r io.Reader) (s Story, err error) {
	d := json.NewDecoder(r)
	if err = d.Decode(&s); err != nil {
		return nil, err
	}
	return s, nil
}
