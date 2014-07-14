package main

import (
	"github.com/go-martini/martini"
	"html/template"
	"io/ioutil"
	"net/http"
	"log"
)

func main() {
	m := martini.Classic()
	m.Get("/", Index)
	m.Get("/cameron-bio", Bio)
	m.Run()
}

type TemplateData struct {
	Title string
}

func Index(w http.ResponseWriter) {

	// head template
	head_file := `views/head.gotmpl`
	head_txt, err := ioutil.ReadFile(head_file)
	if err != nil {
		log.Printf("Error reading %s", head_file)
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// nav template
	nav_file := `views/nav.gotmpl`
	nav_txt, err := ioutil.ReadFile(nav_file)
	if err != nil {
		log.Printf("Error reading %s", nav_file)
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// index template
	tmpl_file := `views/index.gotmpl`
	index_txt, err := ioutil.ReadFile(tmpl_file)
	if err != nil {
		log.Printf("Error reading %s", tmpl_file)
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// put all the templates into one string slice
	index_txt = append(index_txt, head_txt...)
	index_txt = append(index_txt, nav_txt...)
	
	index_data := &TemplateData{ Title: "Home" }
	index_tmpl, err := template.New("index").Parse(string(index_txt))
	if err != nil {
	    log.Printf("Error parsing %s", tmpl_file)
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


func Bio(w http.ResponseWriter) {

	// head template
	head_file := `views/head.gotmpl`
	head_txt, err := ioutil.ReadFile(head_file)
	if err != nil {
		log.Printf("Error reading %s", head_file)
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// nav template
	nav_file := `views/nav.gotmpl`
	nav_txt, err := ioutil.ReadFile(nav_file)
	if err != nil {
		log.Printf("Error reading %s", nav_file)
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// bio template
	tmpl_file := `views/bio.gotmpl`
	bio_txt, err := ioutil.ReadFile(tmpl_file)
	if err != nil {
		log.Printf("Error reading %s", tmpl_file)
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// put all the templates into one string slice
	bio_txt = append(bio_txt, head_txt...)
	bio_txt = append(bio_txt, nav_txt...)

	bio_data := &TemplateData{ Title: "Cameron's Bio" }
	bio_tmpl, err := template.New("bio").Parse(string(bio_txt))
	if err != nil {
	    log.Printf("Error parsing %s", tmpl_file)
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


