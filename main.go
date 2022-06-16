package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/izaakdale/chooseAdventure/jsonStory"
	"github.com/izaakdale/chooseAdventure/storyHandler"
)

func main() {

	jsonFileString := flag.String("story", "gopher.json", "Name of the json file to use as the story")
	htmlFileString := flag.String("template", "template.html", "Name of the html file to use")
	flag.Parse()

	jsonFile, err := os.Open(*jsonFileString)
	if err != nil {
		log.Fatal("Error opening json : " + err.Error())
	}

	htmlBuf, err := os.ReadFile(*htmlFileString)
	if err != nil {
		log.Fatal("Error opening html : " + err.Error())
	}

	story, err := jsonStory.GetAdventureStories(jsonFile)
	if err != nil {
		panic(err)
	}

	htmlTemplate := template.Must(template.New("").Parse(string(htmlBuf)))

	sh := storyHandler.Handler{
		Story: story,
		Html:  htmlTemplate,
	}
	log.Fatal(http.ListenAndServe(":8000", sh))
}
