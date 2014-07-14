package main

import (
	"html/template"
	"net/http"
	"log"
)

func Contact(w http.ResponseWriter) {

	my_views := append(views(), "views/contact.gotmpl")

	views_txt, err := get_templates(my_views)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return		
	}

	contact_data := &TemplateData{ Title: "Contact" }
	contact_tmpl, err := template.New("contact").Parse(string(views_txt))
	if err != nil {
	    log.Printf("Error parsing %s", views_txt)
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = contact_tmpl.Execute(w, contact_data)
	if err != nil {
		log.Printf("Error rendering template %v", contact_data)
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
