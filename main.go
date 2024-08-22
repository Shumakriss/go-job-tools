package main

import (
	"go-job-tools/internal/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/job-post", handlers.JobPostUrlHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
