package main

import (
	"log"
	"net/http"

	"fmt"

	"github.com/gorilla/mux"
	"github.com/middlehut/workshop.k8s/internal/diagnostics"
)

func main() {
	log.Println("hello")

	router := mux.NewRouter()
	router.HandleFunc("/", hello)

	go func() {
		err := http.ListenAndServe(":8080", router)
		if err != nil {
			log.Fatal(err)
		}
	}()

	diags := diagnostics.NewDiagnostics()
	err := http.ListenAndServe(":8585", diags)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, http.StatusText(http.StatusOK))
}
