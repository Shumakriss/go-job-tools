package handlers

import (
	"encoding/json"
	"fmt"
	"go-job-tools/internal/scrape"
	"net/http"
)

type JobPostRequest struct {
	JobPostUrl string `json:"job_post_url" bson:"JobPostUrl"`
}

func JobPostUrlHandler(w http.ResponseWriter, r *http.Request) {
	addCorsHeader(w)

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	mv := validateMethodType(r, "POST")
	if !mv.Valid {
		http.Error(w, mv.Msg, mv.StatusCode)
		return
	}

	ctv := validateContentType(r, "application/json")
	if !ctv.Valid {
		http.Error(w, ctv.Msg, ctv.StatusCode)
		return
	}

	var p JobPostRequest
	decErr := json.NewDecoder(r.Body).Decode(&p)

	if decErr != nil {
		http.Error(w, decErr.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Handling URL: " + p.JobPostUrl)
	s := scrape.NewScraper(p.JobPostUrl)
	sErr := s.Scrape()
	if sErr != nil {
		http.Error(w, sErr.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Fprintln(w, s.JobDescription)
	}
}
