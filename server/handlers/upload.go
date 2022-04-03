package handlers

import (
	"encoding/hex"
	"html/template"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"tempbin/server/db"
	"time"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("server/template/upload.html", "server/template/footer.html")
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		returnError(w, "", "/")
		return
	}

	tmpl.Execute(w, nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {

	// parse multipart with max 10mb in memory
	err := r.ParseMultipartForm(50 << 20)
	if err != nil {
		log.Println("Problem parsing the file " + err.Error())
		returnError(w, "", "/")
		// w.WriteHeader(http.StatusBadRequest)
		return
	}
	f, h, err := r.FormFile("file")
	if err != nil {
		log.Println("Form file problem " + err.Error())
		returnError(w, "", "/")
		return
	}
	// max file size is 50MB
	if h.Size > (50 << 20) {
		returnError(w, "Max file size is 50MB", "/")
		return
	}
	name := strings.ToValidUTF8(h.Filename, "")
	// name := h.Filename
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

	duration, err := strconv.Atoi(r.FormValue("duration"))
	if err != nil {
		log.Println(err)
		duration = 10
	}
	if duration > 120 {
		returnError(w, "Maximum 30 minutes", "/")
	}

	go db.InsertDB(ch, name, url, isProtected, password, isLimit, limit, duration)

	// url is the same as the file name
	tmpFile, err := os.Create("./bucket/" + url)
	defer tmpFile.Close()
	if err != nil {
		ch <- false
		log.Println("Error creating a file to store " + err.Error())
		returnError(w, "", "/")
		return
	}

	_, err = io.Copy(tmpFile, f)
	if err != nil {
		ch <- false
		log.Println("Error copying file " + err.Error())
		returnError(w, "", "/")
		return
	}

	ch <- true
	close(ch)

	tmpl, err := template.ParseFiles("server/template/link.html", "server/template/footer.html")
	if err != nil {
		// w.WriteHeader(http.StatusBadRequest)
		returnError(w, "", "/")
		return
	}
	tmpl.Execute(w, url)
	//w.WriteHeader(200)
	// fmt.Fprint(w, url)
	return
}

func getUUID() string {
	rand.Seed(time.Now().UnixNano())
	p := make([]byte, 4)
	rand.Read(p)
	return hex.EncodeToString(p)
}
