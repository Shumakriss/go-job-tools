package scrape

import (
	"fmt"
	"testing"
)

func TestLevelsSearch(t *testing.T) {
	s := NewScraper("https://www.levels.fyi/jobs")
	sErr := s.Scrape()
	if sErr != nil {
		t.Fatalf("test failed with error: %v", sErr)
	} else if s.SearchResults == "" {
		t.Fatalf("No search results found")
	}
}

func TestLevelsJD(t *testing.T) {
	s := NewScraper("https://www.levels.fyi/jobs?jobId=103126329490055878")
	sErr := s.Scrape()
	if sErr != nil {
		t.Fatalf("test failed with error: %v", sErr)
	} else if s.JobDescription == "" {
		t.Fatalf("No job description found")
	}

	fmt.Println(s.JobDescription)
}
