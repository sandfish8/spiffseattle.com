package main

import (
	"github.com/go-martini/martini"
	"io/ioutil"
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






