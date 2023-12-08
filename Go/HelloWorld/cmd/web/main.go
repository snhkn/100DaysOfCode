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

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/about", handlers.AboutHandler)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
