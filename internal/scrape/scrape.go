package scrape

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
)

type JobPostScraper struct {
	Url            string
	JobDescription string
}

type UnsupportedSite struct{}

func (m *UnsupportedSite) Error() string {
	return "Site requested is unsupported and cannot be scraped"
}

func NewScraper(url string) JobPostScraper {
	return JobPostScraper{Url: url}
}

func (s *JobPostScraper) Scrape() error {

	if strings.Contains(s.Url, "outerjoin.us") {
		return s.OuterJoin()
	} else if strings.Contains(s.Url, "linkedin.com") {
		//return s.LinkedIn()
		return &UnsupportedSite{}
	} else {
		return &UnsupportedSite{}
	}
}

func (s *JobPostScraper) OuterJoin() error {
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

	err := c.Visit(s.Url)

	if err != nil {
		return err
	}

	s.JobDescription = result

	return nil
}

func (s *JobPostScraper) LinkedIn() error {
	result := ""
	c := colly.NewCollector()

	c.OnHTML("body", func(e *colly.HTMLElement) {
		fmt.Println("Found the body")
		fmt.Println(e.Text)
		result += e.Text
	})

	c.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Content-Type", "text/html")
		r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Safari/537.36")
		r.Headers.Set("Cookie", "bcookie=\"v=2&84ef1191-6e02-4f74-8392-2a5ce77551c2\"; bscookie=\"v=1&2024052018540181d38849-141f-4b96-8e8f-b5cbf0f4d8bdAQHU-RCFEh9TVMtas2Xw8z8nXR9el322\"; aam_uuid=86605797631039440083692588765488264664; g_state={\"i_l\":0}; liap=true; JSESSIONID=\"ajax:7227053204767430750\"; timezone=America/New_York; li_theme=light; li_theme_set=app; li_sugr=4da95c07-db76-4e80-b37d-a81e0c31a7cc; dfpfpt=d0de23b1a7b74a588ce52693e64418fe; _gcl_au=1.1.1015063.1716309314; s_fid=198AD9597E70CE06-338398DE4C407094; gpv_pn=developer.linkedin.com%2Fproduct-catalog%2Fplugins; s_ips=1294; s_tp=1360; s_tslv=1720722547737; mbox=PC#24dd4b5df4e04f81b5ae88ef84839f0b.34_0#1736274548|session#4d40dcd4c64c45ee88aa82d33fc43144#1720724408; li_at=AQEDAQRDTEIFyu5xAAABj5dakMUAAAGRf8x0h00AMlAsBzE3CD76T2t9sB5P3uru9Y6nn6nMpMTdxeXclcAvxm_frjCVO5BKMC36EPOGMEIMB4no7Uhovd5PXY30i5QIHKLShgJorRogUTExdtVGUdH0; lang=v=2&lang=en-us; sdsc=22%3A1%2C1724261195742%7EJAPP%2C0xJy8F2JiukoJZXy9ckjwWlSqtqI%3D; AMCVS_14215E3D5995C57C0A495C55%40AdobeOrg=1; fptctx2=taBcrIH61PuCVH7eNCyH0FWPWMZs3CpAZMKmhMiLe%252bFqAiA8vn4FaeCsFGXbYlL%252fDV2l992CS6rQjrgfjYDDAeAcqz7eWaVhmGg%252fPYpbdkudOBzQoF1Mv7pvIiFm3w3hapimUIW0gtrbKCazpphnG8wA0VRXacUhzzBK%252bc3jxYKugNM5QpRik1k9f3P7EHbnPoV0kNCzz%252fkULbw2QgR6vGBSkRXjDhfpXfYRaEjTO9NP3ia8V%252b9%252bPFKVmLollntIWUsg3bdikZg%252bEIy0wz6kUi5sAgY4JnjwhnNCv6DPo8wHC9Ux8C%252bovBOWnk1uL3AapLIQ9y6z1hGZ9GcpsTtPi%252fDI6BdRIfe7E7QIq4PDw0Q%253d; AnalyticsSyncHistory=AQJe-oHfD8UX0AAAAZF2MjzHcO3_C32gjPX8Dy7D_UDfyxeZ-Erff_4gOfNrSUm2QZT7Icgi3g_SMXqeRSRu2w; lms_ads=AQF1jhmyJRaWVQAAAZF2Mj2bmR-mS2OB47unJxO8UIs4ruM1qyRE8rmNrLiZWjf12eU7BG4Pm_oYM42gfiRsW3KXyHrrvXVZ; lms_analytics=AQF1jhmyJRaWVQAAAZF2Mj2bmR-mS2OB47unJxO8UIs4ruM1qyRE8rmNrLiZWjf12eU7BG4Pm_oYM42gfiRsW3KXyHrrvXVZ; AMCV_14215E3D5995C57C0A495C55%40AdobeOrg=-637568504%7CMCIDTS%7C19958%7CMCMID%7C87188636865909059873742943016624866835%7CMCAAMLH-1724946208%7C7%7CMCAAMB-1724946208%7C6G1ynYcLPuiQxYZrsz_pkqfLG9yMXBpb2zX5dvJdYQJzPXImdj0y%7CMCOPTOUT-1724348608s%7CNONE%7CvVersion%7C5.1.1%7CMCCIDH%7C-1112400374; UserMatchHistory=AQLBkCxLh7CPoQAAAZF6xgYxuoLIXE-hrWzGLYVcNZ9OhY0-iHRKQPkCh-ltOqkh7W000b_joRkk5vM8br22J3Ijwr3l6dwk_ZfpuJ9I2g-fWmM2IyTxoXU1KIUXOwcZLj8nFU1XrxNVHJgBYqY4s0l54ZTntXDJ1vgFHmJj-zKtpktPSu2wMVhCN1t4xU6GCBGQ3RGOfbH9H6QgrPBnY7ErIUETnGLLl9XVP9aM1gGFpVpDP2YrxoO8VbUr0ecjoYCr_GU_Hn3lnDKHbCNHa0MJi89kiwRhFcSjp16HZVwYYzJBZmPX4eESfNfnK0tmO32OHJykpjhWntbr_c_N1tkmRS6oMfAMKTTvBGWv3ECxPjYctw; lidc=\"b=TB98:s=T:r=T:a=T:p=T:g=3574:u=1227:x=1:i=1724341684:t=1724368988:v=2:sig=AQG_obw2MPL19rUk5dF8owuQhlRPTgtp")
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "\nError:", err)
	})

	err := c.Visit(s.Url)

	if err != nil {
		fmt.Println("Failed to scrape linkedin.com post")
		return err
	}

	s.JobDescription = result

	return nil
}
