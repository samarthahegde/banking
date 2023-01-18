package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"full_name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func greet(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello, world!!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "Ashish", City: "New Delhi", Zipcode: "110075"},
		{Name: "Rob", City: "New Delhi", Zipcode: "110075"},
	}
	if r.Header.Get("Content-Type") == "application/xml" {

		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)

	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)

		// w.Header().Add("Content-Type", "application/json")
		// json.NewEncoder(w).Encode(customers)

	}

}
