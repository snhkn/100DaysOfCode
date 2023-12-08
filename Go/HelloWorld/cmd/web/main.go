package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/snhkn/100DaysOfCode/Go/HelloWorld/pkg/config"
	"github.com/snhkn/100DaysOfCode/Go/HelloWorld/pkg/handlers"
	"github.com/snhkn/100DaysOfCode/Go/HelloWorld/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tmplCache, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tmplCache
	// when in developer mode set useCache to false for testing purposes
	//otherwise set it to true
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.HomeHandler)
	http.HandleFunc("/about", handlers.Repo.AboutHandler)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
