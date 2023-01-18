package app

import (
	"log"
	"net/http"
)

func Start() {

	// define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	//start serverd
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
