package main

import (
	"log"
	"net/http"
	"text/template"
	"time"
)

const (
	Done    = "Done"
	Waiting = "Waiting"
)

var (
	tasks []*Task
)

func init() {
	tasks = []*Task{
		&Task{Name: "Enviar emails", Status: Waiting, Seconds: 4},
		&Task{Name: "Generar PDFS", Status: Waiting, Seconds: 7},
		&Task{Name: "Buscar/Guardar tweets", Status: Waiting, Seconds: 9},
	}
}

type Task struct {
	Name    string
	Status  string
	Seconds time.Duration
}

func (t *Task) Run() {
	time.Sleep(t.Seconds * time.Second)
	t.Status = Done
	log.Printf("task done: %v", t)
}

func RunTasks() {
	for _, task := range tasks {
		go task.Run()
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, tasks)
}

func StartHandler(w http.ResponseWriter, r *http.Request) {
	go RunTasks()
	http.Redirect(w, r, "localhost:8080/", http.StatusFound)
}

func main() {
	log.Printf("Starting server on :8080")
	http.HandleFunc("/run", StartHandler)
	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":8080", nil)
}
