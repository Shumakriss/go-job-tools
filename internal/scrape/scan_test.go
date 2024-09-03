package scrape

import (
	"fmt"
	"net/http"
	"testing"
)

func TestScan(t *testing.T) {
	s, err := Scan("test", "test")
	if err != nil {
		t.Fatalf("test failed with error: %v", err)
	} else if s.Response == "" {
		t.Fatalf("No job description found")
	}

	fmt.Println(s)
}

func TestFetch(t *testing.T) {

	//postBody, _ := json.Marshal(map[string]any{
	//"cv":                        "test",
	//"jd":                        "test",
	//"coverletter":               "",
	//"coverletter_conversion_id": nil,
	//"conversion_id":             nil,
	//})

	//responseBody := bytes.NewBuffer(postBody)

	//Leverage Go's HTTP Post function to make request
	req, _ := http.NewRequest("GET", "https://api.jobscan.co/v4/scan/36447312", nil)
	//req, _ := http.NewRequest(http.MethodPost, "https://api.jobscan.co/v4/scan", responseBody)
	//req, _ := http.NewRequest("GET", "http://httpbin.org/user-agent", nil)

	//req.Header.Set("X-Xsrf-Token", "eyJpdiI6ImhSUytydHlvWVRIZEtSbk9vYjdOalE9PSIsInZhbHVlIjoiR1Zxb0NaOG9vQzZrb01FamxzUityTXRZQkFZU2hXd2U1UlJBclMzSXFlK04wZ21XRlo5dUJMZ3RmQ3BtNWJHcE82KzlIMFlZMWwyS0lPSUU0LzAvamNYQWQxOHhhSHdJRUdCK0ZnT0xvWGlRaDVnNk8rRzVtdFVRa1JFbmpZZjciLCJtYWMiOiI2OTQwMzUzODZjNDY0ZDg2ZTI2MTk1ZWE0ZDZmMjJiZWNmY2ZhZjI3NzRiODBlYTk3YTJiY2ZmYzNmMjllYzUwIiwidGFnIjoiIn0=")
	//req.Header.Set("Cookie", "ajs_anonymous_id=fdfa3a40-d6f2-43ab-a86b-6664abc2b52e; _gcl_au=1.1.1727714196.1717691086; rl_anonymous_id=RS_ENC_v3_IjhiZTE2YzkxLWMwZmEtNGUzNC04NWZjLWYzZGQ3OWFlNGI3MiI%3D; rl_page_init_referrer=RS_ENC_v3_IiRkaXJlY3Qi; _tt_enable_cookie=1; _ttp=v9llazNBWqnLzuUKl6sxIXqdjpo; hubspotutk=dad739f52cc7edc8aaa38dc35e0669e3; ajs_user_id=4032590; rl_user_id=RS_ENC_v3_IjQwMzI1OTAi; rl_trait=RS_ENC_v3_eyJpZCI6NDAzMjU5MCwiZW1haWwiOiJjaHJpcy5zaHVtYWtlckBnbWFpbC5jb20iLCJmaXJzdE5hbWUiOiJDaHJpc3RvcGhlciIsImxhc3ROYW1lIjoiU2h1bWFrZXIiLCJwbGFuIjoiZG9tYWluIiwiZG9tYWluIjoia2V5c3RvbmVwYXJ0bmVycy5jb20iLCJjcmVhdGVkQXQiOiIyMDI0LTA0LTA5VDE1OjI0OjIxLjAwMFoifQ%3D%3D; remember_web_3dc7a913ef5fd4b890ecabe3487085573e16cf82=eyJpdiI6IkxvU04rdTlvTTJnYVZScHlzOTU5Wnc9PSIsInZhbHVlIjoiOGY5RytuNHdHSWtUVDU1ZzV3ZVVVK2pMRGVRZWYzVnUzKyt3d0NMWUpoNFVFMkZ4K3BhUXFNZWNJR2Q3Qzd6Y0lWSVpCYXV2MmxVWDhvY1l5Z2pCYmVwRUtRUUhOV0NBeEpnck40aGpvTS95SWVTMHZwVWgrT0JtQnBlRDFydmdEUkFIU2NsenY3TXdiWW9EZDg0WS9TcmxHMkFDU2w1QzIrRE5ONEpiYzBNZ0FpbTI3WHN6SWNRK2xQbmlBek1rU3kwY0YyT2s1cnNUenEwTGppdVBwaW1DS25nbVpLYk80QWZ0VzFOTnJvTT0iLCJtYWMiOiI3NGUxMzY4OGI3MzVkNTM1MzY5OGZjN2UzZDU0MjA2OTVjOTgyZDEwN2EyNjM0N2YyMTZhMzM0NmFiMzEzMzU4IiwidGFnIjoiIn0%3D; _fbp=fb.1.1721660162511.739746909782170836; jobscan_utm=%7B%22http_referrer%22%3A%22https%3A%2F%2Fapp.jobscan.co%2F%22%7D; mp_8372d7d978a2c95c7c831351ed0025f5_mixpanel=%7B%22distinct_id%22%3A%20%22%24device%3A19079a8c2582ed-04a088e2a33668-19525637-384000-19079a8c2582ed%22%2C%22%24device_id%22%3A%20%2219079a8c2582ed-04a088e2a33668-19525637-384000-19079a8c2582ed%22%2C%22%24search_engine%22%3A%20%22google%22%2C%22%24initial_referrer%22%3A%20%22https%3A%2F%2Fwww.google.com%2F%22%2C%22%24initial_referring_domain%22%3A%20%22www.google.com%22%2C%22__mps%22%3A%20%7B%7D%2C%22__mpso%22%3A%20%7B%22%24initial_referrer%22%3A%20%22https%3A%2F%2Fwww.google.com%2F%22%2C%22%24initial_referring_domain%22%3A%20%22www.google.com%22%7D%2C%22__mpus%22%3A%20%7B%7D%2C%22__mpa%22%3A%20%7B%7D%2C%22__mpu%22%3A%20%7B%7D%2C%22__mpr%22%3A%20%5B%5D%2C%22__mpap%22%3A%20%5B%5D%7D; _gid=GA1.2.1280677989.1725043366; _gat_segmentGATracker=1; _clck=2qs5h5%7C2%7Cfor%7C0%7C1618; _ce.irv=returning; cebs=1; _ce.clock_event=1; jobscan_ab_cohort=variant1; _CEFT=EgNwlgpg7hAmBcA1MBHAzAMQA4BcDSASmgIIDQAbAOoCiArsQCoDyAmgLQDCAHEA; _ce.clock_data=-10%2C74.104.157.223%2C1%2Cfe0673f2a48d047b912b27e2a0c02f9f%2CChrome%2CUS; __hstc=122684180.dad739f52cc7edc8aaa38dc35e0669e3.1717691087137.1724445577278.1725043367251.29; __hssrc=1; _ga_0SQJM6BJPW=GS1.1.1725043366.31.1.1725043368.58.0.0; _ga_CT69ZFD9NF=GS1.1.1725043366.31.1.1725043368.0.0.0; _ga_0LLQRF33CN=GS1.1.1725043366.31.1.1725043368.0.0.0; _ga=GA1.2.1085078431.1717691086; _uetsid=a6495db066ff11efbb075dca494c1f00; _uetvid=49e96480242111efbab9594993fd1236; cebsp_=3; __hssc=122684180.2.1725043367251; _clsk=na6k8n%7C1725043369195%7C2%7C1%7Co.clarity.ms%2Fcollect; rl_session=RS_ENC_v3_eyJpZCI6MTcyNTA0MzM2NjI0NSwiZXhwaXJlc0F0IjoxNzI1MDQ1MTY5Nzk0LCJ0aW1lb3V0IjoxODAwMDAwLCJhdXRvVHJhY2siOnRydWUsInNlc3Npb25TdGFydCI6ZmFsc2V9; XSRF-TOKEN=eyJpdiI6ImhSUytydHlvWVRIZEtSbk9vYjdOalE9PSIsInZhbHVlIjoiR1Zxb0NaOG9vQzZrb01FamxzUityTXRZQkFZU2hXd2U1UlJBclMzSXFlK04wZ21XRlo5dUJMZ3RmQ3BtNWJHcE82KzlIMFlZMWwyS0lPSUU0LzAvamNYQWQxOHhhSHdJRUdCK0ZnT0xvWGlRaDVnNk8rRzVtdFVRa1JFbmpZZjciLCJtYWMiOiI2OTQwMzUzODZjNDY0ZDg2ZTI2MTk1ZWE0ZDZmMjJiZWNmY2ZhZjI3NzRiODBlYTk3YTJiY2ZmYzNmMjllYzUwIiwidGFnIjoiIn0%3D; jobscan_session=eyJpdiI6IkpXcVpzVG5kOWdXMzdzUEJKbTd0NXc9PSIsInZhbHVlIjoiV25GRGk2L3B6UEFEVTI0MXYxSy9WZGFpb21EaUhUa290eDhjNnc5aWJGaXNxVzEzamJ5WlFLVzBJVUpyMGRLTThIR2x4cnh2NkxXK3BNMFEzdGpBM09zK0VKTUp6NkYyQ0lRTXgyeVFTVWpsN2lJV2dQei85bTMyREgwTXlwQUoiLCJtYWMiOiIyYzc0ZTdkNDlhNWYyYTVjMmMzNWQ4MjdkMzg2MzNlMjY5YWM4ZDBmZGZmYzkwMGY2YTFlNTgwNTFiZDBmMmU1IiwidGFnIjoiIn0%3D; _ce.s=v~3ee343b804577451d40f18c5e909d131bde9c209~lcw~1725043388828~lva~1725043366689~vpv~23~v11ls~e60dc070-3253-11ef-94c0-d7533378b441~v11.fhb~1725043369061~v11.lhb~1725043369061~flvl~%2CViq3FptKR3A%3A6WEuATOY-C8~v11.cs~366246~v11.s~a6878e50-66ff-11ef-ba13-3f2bf47e5fe7~v11.sla~1725043388828~gtrk.la~m0h2aljx~lcw~1725043388829")

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("User-Agent", "PostmanRuntime/7.39.0")
	//req.Header.Set("Accept", "application/json, text/plain, */*")
	//req.Header.Set("Accept-Encoding", "gzip, deflate, br, zstd")
	//req.Header.Set("Accept-Language", "en-US,en;q=0.9,es-US;q=0.8,es;q=0.7")
	//req.Header.Set("Accepts", "application/json")
	//req.Header.Set("Content-Length", "96")
	//req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Origin", "https://app.jobscan.co")
	//req.Header.Set("Priority", "u=1, i")
	//req.Header.Set("Referer", "https://app.jobscan.co/")
	//req.Header.Set("Sec-Ch-Ua", "\"Not/A)Brand\";v=\"8\", \"Chromium\";v=\"126\", \"Google Chrome\";v=\"126\"")
	//req.Header.Set("Sec-Ch-Ua-Mobile", "?0")
	//req.Header.Set("Sec-Ch-Ua-Platform", "macOS")
	//req.Header.Set("Sec-Fetch-Dest", "empty")
	//req.Header.Set("Sec-Fetch-Mode", "cors")
	//req.Header.Set("Sec-Fetch-Site", "same-site")
	//req.Header.Set("X-Spa-Version", "2024-07-03-a-S")
	//req.Header.Set("Access-Control-Allow-Origin", "*")

	fmt.Println(req)
	fmt.Println(req.Host)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println("Error on response")
		t.Fatal(err)
	} else {
		fmt.Println("No error on response")
		fmt.Println(resp.StatusCode)
		fmt.Println(resp.Body)
	}

}
