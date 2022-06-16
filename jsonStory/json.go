package jsonStory

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type Story map[string]AdventureArc

type AdventureArc struct {
	Title     string             `json:"title"`
	Paragraph []string           `json:"story"`
	Options   []AdventureOptions `json:"options"`
}

type AdventureOptions struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func GetAdventureStories(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

func StoryHandler(s Story, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		arcName := strings.TrimLeft(r.URL.Path, "/")
		arc, found := s[arcName]
		if found {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(arc)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}
