package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Number struct {
	First     float64 `json:"first"`
	Second    float64 `json:"second"`
	Operation string  `json:"operation"`
}

type Answer struct {
	Reply float64 `json:"answer"`
}

func calculate(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var numbers Number
	_ = json.NewDecoder(r.Body).Decode(&numbers)
	var answer float64
	if numbers.Operation == "ADD" {
		answer = numbers.First + numbers.Second
	} else if numbers.Operation == "SUB" {
		answer = numbers.First - numbers.Second
	} else if numbers.Operation == "MUL" {
		answer = numbers.First * numbers.Second
	} else if numbers.Operation == "DIV" {
		answer = numbers.First / numbers.Second
	}
	answerJson := Answer{Reply: answer}
	_ = json.NewEncoder(rw).Encode(answerJson)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/calculate", calculate).Methods("POST")
	r.Handle("/", http.FileServer(http.Dir(".")))
	log.Fatal(http.ListenAndServe(":8000", r))

}
