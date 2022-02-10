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

type Err struct {
	Reply string `json:"answer"`
}

func calculate(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	var numbers Number
	_ = json.NewDecoder(r.Body).Decode(&numbers)
	var answer float64
	switch numbers.Operation {
	case "ADD":
		answer = numbers.First + numbers.Second
	case "SUB":
		answer = numbers.First - numbers.Second
	case "MUL":
		answer = numbers.First * numbers.Second
	case "DIV":
		if numbers.Second == 0 {
			divZero := Err{Reply: "Can't divide by zero :("}
			_ = json.NewEncoder(rw).Encode(divZero)
			return
		}
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
