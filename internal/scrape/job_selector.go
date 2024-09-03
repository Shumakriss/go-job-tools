package scrape

import (
	"strings"
)

type UnsupportedJobSite struct{}

func (m *UnsupportedJobSite) Error() string {
	return "Site requested is unsupported and cannot be scraped"
}

func ScrapeJob(url string) (string, error) {

	var results string
	var err error

	if strings.Contains(url, "outerjoin.us") {
		results, err = OuterJoinJD(url)
	} else if strings.Contains(url, "linkedin.com") {
		//return s.LinkedIn()
		return "", &UnsupportedJobSite{}
	} else if strings.Contains(url, "https://www.levels.fyi/jobs?jobId=") {
		results, err = LevelsFyiJD(url)
	} else {
		return "", &UnsupportedJobSite{}
	}

	return results, err
}
