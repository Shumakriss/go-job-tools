package scrape

import (
	"strings"
)

type UnsupportedSearchSite struct{}

func (m *UnsupportedSearchSite) Error() string {
	return "Site requested is unsupported and cannot be scraped"
}

func ScrapeSearch(url string) ([]string, error) {

	var res []string
	var err error

	if strings.Contains(url, "https://www.levels.fyi/jobs") {
		res, err = LevelsFyiSearch(url)
	} else {
		err = &UnsupportedSearchSite{}
	}

	if err != nil {
		return res, err
	}

	return res, nil
}
