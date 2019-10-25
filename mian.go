package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lucasdasilvabarbosa/rest-api/models"
	"log"
	"net/http"
)

var people []models.Person

func main() {
	people = append(people, models.Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &models.Address{City: "City X", State: "State X"}})
	router := mux.NewRouter()
	router.HandleFunc("/contato", GetPeople).Methods("GET")
	router.HandleFunc("/contato/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/contato/{id}", CreatPerson).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func CreatPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person models.Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, request *http.Request) {

}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	print("chegou")
	json.NewEncoder(w).Encode(people)
}
