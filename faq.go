package main

import (
	"html/template"
	"net/http"
	"log"
)

func Faq(w http.ResponseWriter) {

	my_views := append(views(), "views/faq.gotmpl")

	views_txt, err := get_templates(my_views)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return		
	}

	faq_data := &TemplateData{ Title: "FAQ" }
	faq_tmpl, err := template.New("faq").Parse(string(views_txt))
	if err != nil {
	    log.Printf("Error parsing %s", views_txt)
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = faq_tmpl.Execute(w, faq_data)
	if err != nil {
		log.Printf("Error rendering template %v", faq_data)
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
