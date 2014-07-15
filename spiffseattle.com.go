package main

import (
	"github.com/go-martini/martini"
	"io/ioutil"
	"net/http"
	"html/template"
	"log"
	"fmt"
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

func Bio(w http.ResponseWriter) {

	name := "bio"
	data := &TemplateData{ Title: "Cameron's Bio" }
	my_views := append(views(), fmt.Sprintf("views/%s.gotmpl", name))

	err := render(name, data, my_views, w) 
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)		
	}
}

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

// get_templates accepts a slice of views files and return a concatenated string
// of the content
func get_templates(tmpls []string) ([]byte, error) {
	tmpl_txt := []byte{}
	for _, file := range tmpls {
		txt, err := ioutil.ReadFile(file)
		if err != nil {
			return []byte{}, err
		}		
		tmpl_txt = append(tmpl_txt, txt...)
	}
	return tmpl_txt, nil
}






