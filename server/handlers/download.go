package handlers

import (
	"html/template"
	"io"
	"net/http"
	"os"
	"tempbin/server/db"

	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

type FileInfo struct {
	Name, Url            string
	IsLimit, IsProtected bool
	Limit                int
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	// Get file name
	// Check if password protected
	// Check if download limit
	// fmt.Print(chi.URLParam(r, "url"))
	url := chi.URLParam(r, "url")
	name, isProtected, isLimit, _, limit, err := db.GetData(url)
	info := FileInfo{Name: name, Url: url, IsLimit: isLimit, IsProtected: isProtected, Limit: limit}
	if err != nil {
		// fmt.Fprint(w, err)
		w.WriteHeader(http.StatusBadRequest)
	}

	tmpl, err := template.ParseFiles("server/template/download.html")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// w.Write([]byte(err.Error()))
	}
	tmpl.Execute(w, info)
}

func Download(w http.ResponseWriter, r *http.Request) {

	url := chi.URLParam(r, "url")

	name, isProtected, isLimit, password, limit, err := db.GetData(url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	pw := r.FormValue("password")
	if isProtected {
		err = bcrypt.CompareHashAndPassword([]byte(password), []byte(pw))
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	if isLimit {
		if limit < 1 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//reduce limit by 1
		db.ReduceLimit(url)
	}

	// change file name to actual file
	f, err := os.Open("./bucket/" + url)
	if err != nil {
		// panic(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer f.Close()

	// change filename to filename
	w.Header().Set("Content-Disposition", "attachment; filename="+name)
	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
	io.Copy(w, f)
	w.WriteHeader(200)
	return
}