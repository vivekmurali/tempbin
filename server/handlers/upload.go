package handlers

import (
	"encoding/hex"
	"fmt"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"tempbin/server/db"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("server/template/upload.html")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tmpl.Execute(w, nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// parse multipart with max 10mb
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println("Problem parsing the file " + err.Error())
		log.Println(r.Header)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	f, h, err := r.FormFile("file")
	if err != nil {
		log.Println("Form file problem " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	name := h.Filename
	password := r.FormValue("password")
	isProtected := false
	if len(password) > 0 {
		isProtected = true
	}

	isLimit := true

	limit, err := strconv.Atoi(r.FormValue("limit"))
	if err != nil {
		// log.Println(err)
		limit = 0
		isLimit = false
	}

	url := getUUID()
	ch := make(chan bool)
	// errch := make(chan error)

	// ch, name, url, is_protected, password, is_limit, limit
	// insertDB() tx, with channel to rollback
	// go db.InsertDB(ch, errch, name, url, isProtected, password, isLimit, limit)
	go db.InsertDB(ch, name, url, isProtected, password, isLimit, limit)

	// url is the same as the file name
	tmpFile, err := os.Create("./bucket/" + url)
	defer tmpFile.Close()
	if err != nil {
		ch <- false
		log.Println("Error creating a file to store " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = io.Copy(tmpFile, f)
	if err != nil {
		ch <- false
		log.Println("Error copying file " + err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// err = <-errch
	// if err != nil {
	// 	ch <- false
	// 	log.Println("Error copying file " + err.Error())
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	ch <- true
	close(ch)

	w.WriteHeader(200)
	fmt.Fprint(w, url)
	return
}

func getUUID() string {
	rand.Seed(time.Now().UnixNano())
	p := make([]byte, 4)
	rand.Read(p)
	return hex.EncodeToString(p)
}
