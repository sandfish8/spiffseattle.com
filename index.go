package main

import (
	"html/template"
	"net/http"
	"log"
)

func Index(w http.ResponseWriter) {
	
	my_views := append(views(), "views/index.gotmpl")

	views_txt, err := get_templates(my_views)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return		
	}

	// render the page
	index_data := &TemplateData{ Title: "Home" }
	index_tmpl, err := template.New("index").Parse(string(views_txt))
	if err != nil {
	    log.Printf("Error parsing %s", views_txt)
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = index_tmpl.Execute(w, index_data)
	if err != nil {
		log.Printf("Error rendering template %v", index_data)
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
