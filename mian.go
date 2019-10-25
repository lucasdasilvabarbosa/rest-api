package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/contato", GetPeople).Methods("GET")
	router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", CreatPerson).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

//var people []Person
//
//people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})

func CreatPerson(w http.ResponseWriter, request *http.Request) {

}

func GetPerson(w http.ResponseWriter, request *http.Request) {

}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	print("chegou")
}
