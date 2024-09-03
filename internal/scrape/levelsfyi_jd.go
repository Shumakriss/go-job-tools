package scrape

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
)

func LevelsFyiJD(url string) (string, error) {
	//fmt.Println("Scraping Levels FYI Search")
	result := ""

	c := colly.NewCollector()

	c.OnHTML("script", func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr("id"), "__NEXT_DATA__") {
			result += e.Text
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	err := c.Visit(url)

	if err != nil {
		return "", err
	}

	return result, nil
}
