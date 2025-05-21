package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	// check to see if we already have the template in our cache
	_, inMap := templateCache[t]

	if !inMap {
		// need to create the template
		log.Println("creating template and adding to cache")
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the template in the cache
		log.Println("using cached template")
	}

	tmpl = templateCache[t]
	_ = tmpl.Execute(w, nil)
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// parse the templates
	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	// add template to cache
	templateCache[t] = tmpl

	return nil
}
