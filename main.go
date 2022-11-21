package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type event struct {
	ID          string `json:"ID"`
	Name        string `json: "Name"`
	Description string `json:Description"`
}

type MostEvent []event

var events = MostEvent{
	{
		ID:          "2",
		Name:        "ALi",
		Description: "Im coding",
	},
	{
		ID:          "3",
		Name:        "mohamed",
		Description: "doing well",
	},
}

// my HOme Page
func mainPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "my homePage is here.... ")

}

// HERE IS FUNCTION TO  CREATE NEW EVENT

func createEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent event
	reqBody, err := ioutil.ReadAll(r.Body) // using ioutil to convert  to  slice
	if err != nil {
		fmt.Fprintf(w, "Please Enter your data ")
	}

	json.Unmarshal(reqBody, &newEvent)
	events = append(events, newEvent)
	w.WriteHeader(http.StatusCreated)

	res := json.NewEncoder(w).Encode(newEvent)
	fmt.Println(" THe data has been created successfully", res)

}

// A FUNCTION TO  GET  ONE   EVENT
func getOneEvent(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]

	for _, oneEvent := range events {
		if oneEvent.ID == eventID {
			json.NewEncoder(w).Encode(oneEvent)
		}
	}

}

// A FUNCTION TO  GET  ALL  THE EVENTS

func getALlEvent(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(events)

}

// A FUNCTION  TO  UPDATE A EVENT FOR ONE  EVENT

func updateOneEVENT(w http.ResponseWriter, r *http.Request) {
	eventID := mux.Vars(r)["id"]
	var updatedEvent event

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "enter  data to be updated ")
	}
	json.Unmarshal(reqBody, &updatedEvent)

	for i, oneEvent := range events {
		if oneEvent.ID == eventID {
			oneEvent.Name = updatedEvent.Name
			oneEvent.Description = updatedEvent.Description
			events = append(events[:i], oneEvent)
			json.NewEncoder(w).Encode(oneEvent)
		}
	}

}

// A FUNCTION TO  DELETE AN EVENT

func deleteEvent(w http.ResponseWriter, r *http.Request) {

	eventID := mux.Vars(r)["id"]

	for i, oneEvent := range events {
		if oneEvent.ID == eventID {
			events = append(events[:i], events[i+1:]...)
			fmt.Fprintf(w, " THE EVENT WITH ID %v  has been deleted  successfully", eventID)
		}
	}

}

// our main function
func main() {
	// initEvents()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", mainPage)
	router.HandleFunc("/event", createEvent).Methods("POST")
	router.HandleFunc("/events/{id}", getOneEvent).Methods("GET")
	router.HandleFunc("/events", getALlEvent).Methods("GET")
	router.HandleFunc("/events/{id}", updateOneEVENT).Methods("PATCH")
	router.HandleFunc("/events/{id}", deleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))

}
