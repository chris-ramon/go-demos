package main

import (
	"net/http"
	"encoding/json"
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"io/ioutil"
)

var db, _ = gorm.Open("postgres", "user=chris dbname=higo sslmode=disable")


type Event struct {
	Id int64
	Name string
}

func availableEvents() (resources *gorm.DB) {
	//events = append(events, Event{Name: "Godsmack Concert"})
	//events = append(events, Event{Name: "Linkin Park Concert"})
	var events []*Event
	resources = db.Find(&events)
	return resources
}

func eventCreator(request *http.Request) (event Event, error error) {
	body, _ := ioutil.ReadAll(request.Body)
	data := make(map[string])
	json.Unmarshal(body, &data)
	log.Println(data["name"])
	return Event{Name: "abc"}, nil
}

func eventsHandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		log.Println("GET")
		events, _ := json.Marshal(availableEvents())
		response.Write(events)
	} else if request.Method == "POST" {
		log.Println("POST")
		event, _ := eventCreator(request)
		eventJson, _ := json.Marshal(event)
		response.Write(eventJson)
	}
}

func main() {

	http.HandleFunc("/events", eventsHandler)
	http.ListenAndServe(":8080", nil)
}


