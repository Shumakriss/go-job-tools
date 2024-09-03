package scrape

import (
	"testing"
)

func TestOuterJoinJD(t *testing.T) {
	// TODO: replace test url with snapshot of HTML file in repo
	s, err := OuterJoinJD("https://outerjoin.us/remote-jobs/senior-data-engineer-1-at-loom")
	if err != nil {
		t.Fatalf("test failed with error: %v", err)
	} else if s == "" {
		t.Fatalf("No job description found")
	}

}
