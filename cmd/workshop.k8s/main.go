package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("hello")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
