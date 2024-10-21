package web

import "html/template"

func LoadTemplates() (*template.Template, error) {
	file := "templates/*.html"

	template, err := template.ParseGlob(file)
	if err != nil {
		panic(err)
	}

	return template, nil
}
