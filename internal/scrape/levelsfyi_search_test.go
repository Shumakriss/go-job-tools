package scrape

import (
	"fmt"
	"testing"
)

func TestLevelsSearch(t *testing.T) {
	s, err := LevelsFyiSearch("https://www.levels.fyi/jobs")
	if err != nil {
		t.Fatalf("test failed with error: %v", err)
	} else if len(s) == 0 {
		t.Fatalf("No job description found")
	}

	fmt.Println(s)
}
