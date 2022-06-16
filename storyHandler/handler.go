package storyHandler

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/izaakdale/chooseAdventure/jsonStory"
)

type Handler struct {
	Story jsonStory.Story
	Html  *template.Template
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	arcName := strings.TrimLeft(r.URL.Path, "/")
	arc, found := h.Story[arcName]
	if found {
		err := h.Html.Execute(w, arc)
		if err != nil {
			log.Fatal("Error executing : " + err.Error())
		}
	} else {
		err := h.Html.Execute(w, h.Story["intro"])
		if err != nil {
			log.Fatal("Error executing : " + err.Error())
		}
	}
}
