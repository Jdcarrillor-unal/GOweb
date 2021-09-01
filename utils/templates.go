package utils

import (
	"api/config"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.New("t").ParseGlob(config.DirTemplate()))
var errorTemplate = template.Must(template.ParseFiles(config.DirTempalteError()))

func RenderErrorTemplate(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	errorTemplate.Execute(w, nil)
}

func RenderTemplate(w http.ResponseWriter, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html")
	err := templates.ExecuteTemplate(w, name, data)

	if err != nil {
		log.Println(err)
		RenderErrorTemplate(w, http.StatusInternalServerError)
	}
}
