package main

import (
	"go-job-tools/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/job", handlers.Job)
	http.HandleFunc("/search", handlers.Search)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
