package main

import (
	"html/template"
	"net/http"
	"log"
	"fmt"
)

func Contact(w http.ResponseWriter) {

	name := "contact"
	data := &TemplateData{ Title: "Contact" }
	my_views := append(views(), fmt.Sprintf("views/%s.gotmpl", name))

	tmpl, err := template.ParseFiles(my_views...)
	if err != nil {
		log.Println("Error parsing template files")
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Printf("Error rendering template %v", data)
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
