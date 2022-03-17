package main

import (
	"fmt"
	"log"
	"os/exec"
	"tempbin/server/db"
)

func worker() {

	files, err := db.GetToDelete()
	if err != nil {
		log.Println(err)
	}
	for _, v := range files {
		fmt.Printf("v = %+v\n", v)
		cmd := exec.Command("rm", "./bucket/"+v)
		err := cmd.Run()
		if err != nil {
			log.Println(err)
		}
	}
}
