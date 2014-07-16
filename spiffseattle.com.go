package main

import (
	"fmt"
	"github.com/go-martini/martini"
	"html/template"
	"log"
	"net/http"
)

func main() {
	m := martini.Classic()
	m.Get("/", Index)
	m.Get("/cameron-bio", Bio)
	m.Get("/faq", Faq)
	m.Get("/contact", Contact)
	m.Run()
}

type TemplateData struct {
	Title string
}

// views provides a slice of views
// that most of the resource handlers will end up loading
func views() []string {
	return []string{"views/head.gotmpl",
		"views/nav.gotmpl",
		"views/footer.gotmpl",
		"views/bottom.gotmpl"}
}

func Index(w http.ResponseWriter) {

	name := "index"
	data := &TemplateData{Title: "Home"}
	my_views := append(views(), fmt.Sprintf("views/%s.gotmpl", name))

	err := render(name, data, my_views, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Bio(w http.ResponseWriter) {

	name := "bio"
	data := &TemplateData{Title: "Cameron's Bio"}
	my_views := append(views(), fmt.Sprintf("views/%s.gotmpl", name))

	err := render(name, data, my_views, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Faq(w http.ResponseWriter) {

	name := "faq"
	data := &TemplateData{Title: "FAQ"}
	my_views := append(views(), fmt.Sprintf("views/%s.gotmpl", name))

	err := render(name, data, my_views, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func Contact(w http.ResponseWriter) {

	name := "contact"
	data := &TemplateData{Title: "Contact"}
	my_views := append(views(), fmt.Sprintf("views/%s.gotmpl", name))

	err := render(name, data, my_views, w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// render renders the template and writes the response to the http.ResponseWriter
func render(name string, data *TemplateData, views []string, w http.ResponseWriter) error {
	tmpl, err := template.ParseFiles(views...)
	if err != nil {
		log.Println("Error parsing template files")
		log.Println(err)
		return err
	}

	err = tmpl.ExecuteTemplate(w, name, data)
	if err != nil {
		log.Printf("Error rendering template %v", data)
		log.Println(err.Error())
		return err
	}
	return nil
}
