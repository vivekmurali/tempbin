package main

import (
	"fmt"
	"log"
	"os/exec"
	"tempbin/server/db"
	"tempbin/server/handlers"
)

func worker() {

	files, count, err := db.GetToDelete()
	if err != nil {
		log.Println(err)
	}

	handlers.NumFiles.WithLabelValues().Set(count)

	for _, v := range files {
		fmt.Printf("v = %+v\n", v)
		cmd := exec.Command("rm", "./bucket/"+v)
		err := cmd.Run()
		if err != nil {
			log.Println(err)
		}
	}
}
