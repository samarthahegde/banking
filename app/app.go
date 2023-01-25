package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/samarthahegde/banking/domain"
	"github.com/samarthahegde/banking/service"
)

func Start() {

	// mux := http.NewServeMux()
	router := mux.NewRouter()

	//wiring

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	// router.HandleFunc("/greet", greet).Methods(http.MethodGet)
	// router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	// router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	//start serverd
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

// func createCustomer(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Post request recieved")
// }

// func getCustomer(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Fprint(w, vars["customer_id"])

// }
