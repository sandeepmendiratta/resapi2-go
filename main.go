package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	router := NewRouter()

	fmt.Println("Running server!")
	log.Fatal(http.ListenAndServe(":8080", router))
}
