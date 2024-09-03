package scrape

import (
	"fmt"
	"testing"
)

func TestLevelsJD(t *testing.T) {
	// TODO: replace test url with snapshot of HTML file in repo
	s, err := LevelsFyiJD("https://www.levels.fyi/jobs?jobId=103126329490055878")
	if err != nil {
		t.Fatalf("test failed with error: %v", err)
	} else if s == "" {
		t.Fatalf("No job description found")
	}

	fmt.Println(s)
}
