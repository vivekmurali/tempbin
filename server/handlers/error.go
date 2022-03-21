package handlers

import (
	"html/template"
	"net/http"
)

type ErrorData struct {
	Err, Link string
}

func returnError(w http.ResponseWriter, errorString, link string) {

	data := ErrorData{Err: errorString, Link: link}
	tmpl, _ := template.ParseFiles("server/template/error.html")
	tmpl.Execute(w, data)
	w.WriteHeader(http.StatusBadRequest)
}
