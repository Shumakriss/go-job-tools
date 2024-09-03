package handlers

import (
	"encoding/json"
	"fmt"
	"go-job-tools/internal/scrape"
	"net/http"
)

type JobRequest struct {
	Url string `json:"url" bson:"Url"`
}

func Job(w http.ResponseWriter, r *http.Request) {
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

	var p JobRequest
	decErr := json.NewDecoder(r.Body).Decode(&p)

	if decErr != nil {
		http.Error(w, decErr.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Handling URL: " + p.Url)
	jobDesc, sErr := scrape.ScrapeJob(p.Url)

	if sErr != nil {
		http.Error(w, sErr.Error(), http.StatusInternalServerError)
		return
	} else {
		if jobDesc != "" {
			fmt.Fprintln(w, jobDesc)
		}
	}
}
