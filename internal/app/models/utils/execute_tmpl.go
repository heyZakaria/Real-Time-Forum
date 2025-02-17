package utils

import (
	"html/template"
	"net/http"
)

func ExecuteTemplate(w http.ResponseWriter, pages []string, data any) {

	tmpl, err := template.ParseFiles(pages...)
	if err != nil {
		http.Error(w, "Internl Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)

	if err != nil {
		http.Error(w, "Internl Server Error", http.StatusInternalServerError)
		return
	}
}
