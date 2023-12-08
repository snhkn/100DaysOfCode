package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// RenderTemplate renders templates using html/template
func RenderTemplateTest(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing the template:", err)
		return
	}

}

// tmplCache is template cache to store templates in memory
var tmplCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tmplPtr *template.Template
	var err error

	//check if we have already the template in cache
	_, inMap := tmplCache[tmpl]

	// if the template is not in our cache
	if !inMap {
		//we need to create the template
		//we need to parse the tamplate and add it to our cache
		log.Println("Creating parsed template and adding to cache")
		err = createTemplateCache(tmpl)
		if err != nil {
			log.Println(err)
		}
	} else {
		//we have the template in our cache
		log.Println("Using cached template")
	}

	tmplPtr = tmplCache[tmpl]
	err = tmplPtr.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	//Parse the templates
	parsedTmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	// add template to cache(map)
	tmplCache[t] = parsedTmpl
	return nil
}
