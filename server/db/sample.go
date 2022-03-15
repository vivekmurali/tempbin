package db

import (
	"context"
	"fmt"
	"log"
)

func HelloWorld() {
	var greeting string
	err := Env.DB.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(greeting)
}
