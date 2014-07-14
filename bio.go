package main

import (
	"html/template"
	"net/http"
	"log"
)

func Bio(w http.ResponseWriter) {

	views := []string{"views/head.gotmpl", 
	                  "views/nav.gotmpl",
	                  "views/bio.gotmpl"}

	views_txt, err := get_templates(views)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return		
	}

	bio_data := &TemplateData{ Title: "Cameron's Bio" }
	bio_tmpl, err := template.New("bio").Parse(string(views_txt))
	if err != nil {
	    log.Printf("Error parsing %s", views_txt)
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = bio_tmpl.Execute(w, bio_data)
	if err != nil {
		log.Printf("Error rendering template %v", bio_data)
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
