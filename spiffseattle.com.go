package main

import (
	"github.com/go-martini/martini"
	"io/ioutil"
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






