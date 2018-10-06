package main

import (
	"log"
	"net/http"

	"fmt"

	"os"

	"github.com/gorilla/mux"
	"github.com/middlehut/workshop.k8s/internal/diagnostics"
)

func main() {
	log.Println("starting server...")

	blport := os.Getenv("PORT")
	if len(blport) == 0 {
		log.Fatal("PORT is missing")
	}

	diagPort := os.Getenv("DIAG_PORT")
	if len(diagPort) == 0 {
		log.Fatal("DIAG_PORT is missing")
	}

	router := mux.NewRouter()
	router.HandleFunc("/", hello)

	go func() {
		log.Println("Application server is preparing to handle connections...")
		err := http.ListenAndServe(":"+blport, router)
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Diagnostics server is preparing to handle connections...")
	diags := diagnostics.NewDiagnostics()
	err := http.ListenAndServe(":"+diagPort, diags)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("hello handler was called")
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}
