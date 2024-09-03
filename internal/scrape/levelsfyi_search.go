package scrape

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
)

func LevelsFyiSearch(url string) ([]string, error) {
	//fmt.Println("Scraping Levels FYI Search")
	var res []string

	c := colly.NewCollector()

	c.OnHTML("div", func(e *colly.HTMLElement) {
		//fmt.Println("Found div")
		if e.Attr("role") == "button" && strings.Contains(e.Attr("class"), "company-jobs-preview-card_container") {
			//fmt.Println("It's a company preview")
			e.ForEach("div",
				func(i int, child *colly.HTMLElement) {
					//fmt.Println("Found a job link")
					if strings.Contains(child.Attr("class"), "company-jobs-preview-card_companyJobsContainer") {
						child.ForEach("a",
							func(i int, cc *colly.HTMLElement) {
								//fmt.Println(cc.Attr("href"))
								res = append(res, "https://www.levels.fyi"+cc.Attr("href"))
							})
					}
				})
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	err := c.Visit(url)

	if err != nil {
		return res, err
	}

	return res, nil
}
