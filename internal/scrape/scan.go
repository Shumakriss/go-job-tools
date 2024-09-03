package scrape

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
)

// const JobscanUrl = "https://api.jobscan.co/v4/scan"
const JobscanUrl = "https://api.jobscan.co/v4/scan/36447312"

type JobScanResponse struct {
	Response string
}

type JobScanRequest struct {
	Cv                      string `json:"cv" bson:"Cv"`
	Jd                      string `json:"jd" bson:"Jd"`
	CoverLetter             string `json:"coverletter" bson:"Coverletter"`
	ConversionId            string `json:"conversion_id" bson:"ConversionId"`
	CoverLetterConversionId string `json:"coverletter_conversion_id" bson:"CoverletterConversionId"`
}

func Scan(resume string, jobDescription string) (JobScanResponse, error) {
	//fmt.Println("Scraping Levels FYI Search")
	var result JobScanResponse

	c := colly.NewCollector()
	c.UserAgent = "PostmanRuntime/7.39.0"

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", JobscanUrl)
		r.Headers.Set("Accept", "*/*")
		r.Headers.Set("Accept-Encoding", "gzip, deflate, br")
		r.Headers.Set("Connection", "keep-alive")
		//r.Headers.Set("Accept-Language", "en-US,en;q=0.9,es-US;q=0.8,es;q=0.7")
		//r.Headers.Set("Accepts", "application/json")
		//r.Headers.Set("Content-Length", "96")
		r.Headers.Set("Content-Type", "application/json")
		//r.Headers.Set("Cookie", "ajs_anonymous_id=fdfa3a40-d6f2-43ab-a86b-6664abc2b52e; _gcl_au=1.1.1727714196.1717691086; rl_anonymous_id=RS_ENC_v3_IjhiZTE2YzkxLWMwZmEtNGUzNC04NWZjLWYzZGQ3OWFlNGI3MiI%3D; rl_page_init_referrer=RS_ENC_v3_IiRkaXJlY3Qi; _tt_enable_cookie=1; _ttp=v9llazNBWqnLzuUKl6sxIXqdjpo; hubspotutk=dad739f52cc7edc8aaa38dc35e0669e3; ajs_user_id=4032590; rl_user_id=RS_ENC_v3_IjQwMzI1OTAi; rl_trait=RS_ENC_v3_eyJpZCI6NDAzMjU5MCwiZW1haWwiOiJjaHJpcy5zaHVtYWtlckBnbWFpbC5jb20iLCJmaXJzdE5hbWUiOiJDaHJpc3RvcGhlciIsImxhc3ROYW1lIjoiU2h1bWFrZXIiLCJwbGFuIjoiZG9tYWluIiwiZG9tYWluIjoia2V5c3RvbmVwYXJ0bmVycy5jb20iLCJjcmVhdGVkQXQiOiIyMDI0LTA0LTA5VDE1OjI0OjIxLjAwMFoifQ%3D%3D; remember_web_3dc7a913ef5fd4b890ecabe3487085573e16cf82=eyJpdiI6IkxvU04rdTlvTTJnYVZScHlzOTU5Wnc9PSIsInZhbHVlIjoiOGY5RytuNHdHSWtUVDU1ZzV3ZVVVK2pMRGVRZWYzVnUzKyt3d0NMWUpoNFVFMkZ4K3BhUXFNZWNJR2Q3Qzd6Y0lWSVpCYXV2MmxVWDhvY1l5Z2pCYmVwRUtRUUhOV0NBeEpnck40aGpvTS95SWVTMHZwVWgrT0JtQnBlRDFydmdEUkFIU2NsenY3TXdiWW9EZDg0WS9TcmxHMkFDU2w1QzIrRE5ONEpiYzBNZ0FpbTI3WHN6SWNRK2xQbmlBek1rU3kwY0YyT2s1cnNUenEwTGppdVBwaW1DS25nbVpLYk80QWZ0VzFOTnJvTT0iLCJtYWMiOiI3NGUxMzY4OGI3MzVkNTM1MzY5OGZjN2UzZDU0MjA2OTVjOTgyZDEwN2EyNjM0N2YyMTZhMzM0NmFiMzEzMzU4IiwidGFnIjoiIn0%3D; mp_8372d7d978a2c95c7c831351ed0025f5_mixpanel=%7B%22distinct_id%22%3A%20%22%24device%3A19079a8c2582ed-04a088e2a33668-19525637-384000-19079a8c2582ed%22%2C%22%24device_id%22%3A%20%2219079a8c2582ed-04a088e2a33668-19525637-384000-19079a8c2582ed%22%2C%22%24search_engine%22%3A%20%22google%22%2C%22%24initial_referrer%22%3A%20%22https%3A%2F%2Fwww.google.com%2F%22%2C%22%24initial_referring_domain%22%3A%20%22www.google.com%22%2C%22__mps%22%3A%20%7B%7D%2C%22__mpso%22%3A%20%7B%22%24initial_referrer%22%3A%20%22https%3A%2F%2Fwww.google.com%2F%22%2C%22%24initial_referring_domain%22%3A%20%22www.google.com%22%7D%2C%22__mpus%22%3A%20%7B%7D%2C%22__mpa%22%3A%20%7B%7D%2C%22__mpu%22%3A%20%7B%7D%2C%22__mpr%22%3A%20%5B%5D%2C%22__mpap%22%3A%20%5B%5D%7D; _ga_CT69ZFD9NF=GS1.1.1721059195.18.0.1721059195.0.0.0; _ga_0SQJM6BJPW=GS1.1.1721059195.18.0.1721059195.60.0.0; _ga_0LLQRF33CN=GS1.1.1721059195.18.0.1721059195.0.0.0; _ga=GA1.2.1085078431.1717691086; _gid=GA1.2.1439859116.1721059195; _gat_segmentGATracker=1; _ce.irv=returning; cebs=1; _ce.clock_event=1; _clck=2qs5h5%7C2%7Cfnh%7C0%7C1618; _ce.clock_data=-22%2C74.104.157.223%2C1%2C10f9287deaf609ee36fb37783f2b89c0%2CChrome%2CUS; __hstc=122684180.dad739f52cc7edc8aaa38dc35e0669e3.1717691087137.1720624550415.1721059195806.19; __hssrc=1; _uetsid=4766168042c311efba14e1cd99637dc8; _uetvid=49e96480242111efbab9594993fd1236; cebsp_=2; __hssc=122684180.2.1721059195806; _clsk=1goozxm%7C1721059196539%7C2%7C1%7Co.clarity.ms%2Fcollect; rl_session=RS_ENC_v3_eyJpZCI6MTcyMTA1OTE5NTMwMiwiZXhwaXJlc0F0IjoxNzIxMDYwOTk3MzAyLCJ0aW1lb3V0IjoxODAwMDAwLCJhdXRvVHJhY2siOnRydWUsInNlc3Npb25TdGFydCI6ZmFsc2V9; XSRF-TOKEN=eyJpdiI6IkhJY21yZnBiTUNoaVhZY0U1U043V0E9PSIsInZhbHVlIjoiV0h5QzFpaWkwUDlsaUJUTE1EaVVSeVQvODhxYUhGb2hXQjJSYU1iMyt4bmhndE16SWJZNjZ5dFRrNjV3dElJa2prOUdTQ0xIR0IrQ2Zxd2psTGEwQVRUSHNUaTE2WnhjWUM1d3AvYi82T0I0cGdwUkZFRlJXdFNLSGYzUHRhb3IiLCJtYWMiOiJkYjBmMjZiN2Q3ZTRhMmFjYTliYmRhMDUxYTE5NzQxMDgxYTZjYmU1YTc3NGUzMmZiZDI2MjY3MzFhYzkxOTc0IiwidGFnIjoiIn0%3D; jobscan_session=eyJpdiI6IkEzWU9JcGQrZFNBMWs2WHhBVlNPdHc9PSIsInZhbHVlIjoiVjIyY3djWWJJbVB4ZjB5WUJEZk50M1hmWWJNY0VkeDJNd1ZPK21pL3k1TWVhSnlkaGszME5xdlF2Z3FiR0dvWmtVcFlrYmFYZ29PVkU1Rzh0ZHFUeUFyRFRIdEFudnFaMERDaTA3T01nQjVSWmo2WFRtMHpxK0FLbCt6bDA3RHIiLCJtYWMiOiJjYmQyMGMwNWRkYTZkYjYwYjM3Y2IxZmU0ZDE2MGY4ZDY3OWU0MWNmZjRiNDkxNzNiNGQyMjc3NWQxM2NkZTdhIiwidGFnIjoiIn0%3D; _ce.s=v~3ee343b804577451d40f18c5e909d131bde9c209~lcw~1721059217609~lva~1721059195615~vpv~13~v11ls~e60dc070-3253-11ef-94c0-d7533378b441~v11.fhb~1721059196338~v11.lhb~1721059196338~v11.cs~366246~v11.s~47826ac0-42c3-11ef-9b47-f3926c173db7~v11.sla~1721059217625~gtrk.la~lyn67z7j~lcw~1721059217625")
		//r.Headers.Set("Origin", "https://app.jobscan.co")
		//r.Headers.Set("Priority", "u=1, i")
		//r.Headers.Set("Referer", "https://app.jobscan.co/")
		//r.Headers.Set("Sec-Ch-Ua", "\"Not/A)Brand\";v=\"8\", \"Chromium\";v=\"126\", \"Google Chrome\";v=\"126\"")
		//r.Headers.Set("Sec-Ch-Ua-Mobile", "?0")
		//r.Headers.Set("Sec-Ch-Ua-Platform", "macOS")
		//r.Headers.Set("Sec-Fetch-Dest", "empty")
		//r.Headers.Set("Sec-Fetch-Mode", "cors")
		//r.Headers.Set("Sec-Fetch-Site", "same-site")
		//r.Headers.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")
		//r.Headers.Set("X-Spa-Version", "2024-07-03-a-S")
		//r.Headers.Set("X-Xsrf-Token", "eyJpdiI6IkhJY21yZnBiTUNoaVhZY0U1U043V0E9PSIsInZhbHVlIjoiV0h5QzFpaWkwUDlsaUJUTE1EaVVSeVQvODhxYUhGb2hXQjJSYU1iMyt4bmhndE16SWJZNjZ5dFRrNjV3dElJa2prOUdTQ0xIR0IrQ2Zxd2psTGEwQVRUSHNUaTE2WnhjWUM1d3AvYi82T0I0cGdwUkZFRlJXdFNLSGYzUHRhb3IiLCJtYWMiOiJkYjBmMjZiN2Q3ZTRhMmFjYTliYmRhMDUxYTE5NzQxMDgxYTZjYmU1YTc3NGUzMmZiZDI2MjY3MzFhYzkxOTc0IiwidGFnIjoiIn0=")

	})

	c.OnHTML("script", func(e *colly.HTMLElement) {
		if strings.Contains(e.Attr("id"), "__NEXT_DATA__") {
			result.Response += e.Text
		}
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
		fmt.Println(hex.EncodeToString(r.Body))
	})

	payload := JobScanRequest{resume, jobDescription, "", "", ""}
	fmt.Println(payload)
	jsonString, err := json.Marshal(payload)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(jsonString))
	}

	//pErr := c.PostRaw(JobscanUrl, jsonString)
	pErr := c.Visit(JobscanUrl)

	if pErr != nil {
		return result, pErr
	}

	return result, nil
}
