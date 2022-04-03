package db

import (
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// ch, name, url, is_protected, password, is_limit, limit
func InsertDB(ch chan bool, name string, url string, isProtected bool, password string, isLimit bool, limit int, duration int) {
	// func InsertDB(ch chan bool, errch chan error, name string, url string, isProtected bool, password string, isLimit bool, limit int) {

	tx, _ := Env.DB.Begin(context.Background())
	// if err != nil {
	// 	errch <- err
	// }

	pw, _ := bcrypt.GenerateFromPassword([]byte(password), 8)
	// if err != nil {
	// 	errch <- err
	// }

	tag, err := tx.Exec(context.Background(), "insert into files(name, url, is_protected, password, is_limit, \"limit\", duration) values($1, $2, $3, $4, $5, $6, $7)", name, url, isProtected, pw, isLimit, limit, duration)
	if err != nil {
		log.Println(err)
		tx.Rollback(context.Background())
	}
	if !tag.Insert() {
		tx.Rollback(context.Background())
	}

	msg := <-ch
	if msg {
		err = tx.Commit(context.Background())
		if err != nil {
			log.Fatal(err)
		}

	} else {
		tx.Rollback(context.Background())
	}
}
