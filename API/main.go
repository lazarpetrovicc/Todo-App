package main

import (
	"log"
	"net/http"
)

func main() {
	initMongo()

	http.HandleFunc("/todos", todosHandler)
	http.HandleFunc("/todos/", todoHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
