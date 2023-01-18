package app

import (
	"log"
	"net/http"
)

func start() {

	// define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	//start server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
