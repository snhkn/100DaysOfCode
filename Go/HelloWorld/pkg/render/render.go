package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	tmplCache, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}
	// get requested template from cache
	t, ok := tmplCache[tmpl]
	if !ok {
		log.Fatal(err)
	}

	//for debugging
	buf := new(bytes.Buffer)
	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	//get all of the files *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	//range through all files having pattern *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		//parse pages
		tmplSet, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		//get all of the layouts *.layout.tmpl from ./templates
		layouts, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(layouts) > 0 {
			//parse layout tamplates
			tmplSet, err = tmplSet.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = tmplSet
	}

	return myCache, nil

}
