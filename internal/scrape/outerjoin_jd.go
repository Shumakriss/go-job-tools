package scrape

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func OuterJoinJD(url string) (string, error) {
	result := ""
	c := colly.NewCollector()

	c.OnHTML("div .job-description", func(e *colly.HTMLElement) {
		e.ForEach("p",
			func(i int, child *colly.HTMLElement) {
				result += child.Text + "\n"
			})
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
