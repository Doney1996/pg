package main

import (
	"log"
	"person/manage/router"
)

func main() {
	server := router.StepRouter()
	err := server.Run(":8080")
	if err != nil {
		log.Fatal(err, "----- start server fail")
	}
}
