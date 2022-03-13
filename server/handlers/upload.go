package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
)

func Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println("Problem parsing the file " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	f, h, err := r.FormFile("file")
	if err != nil {
		log.Println("Form file problem " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tmpFile, err := os.Create("./bucket/" + h.Filename)
	defer tmpFile.Close()
	if err != nil {
		log.Println("Creating a file to store " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = io.Copy(tmpFile, f)
	if err != nil {
		log.Println("Copying file " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(200)
	return
}
