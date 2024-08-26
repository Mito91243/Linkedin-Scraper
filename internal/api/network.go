package api

import (
	"fmt"
	"io"
	"main/internal/utils"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func GetReq(url string, client *http.Client) ([]byte, int) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("❌ Error creating request:", err)
		return nil, 0
	}
	req.Header.Add("accept", "application/vnd.linkedin.normalized+json+2.1")
	req.Header.Add("accept-encoding", "gzip, deflate, br, zstd")
	req.Header.Add("accept-language", "en-GB,en-US;q=0.9,en;q=0.8")
	req.Header.Add("Csrf-Token", os.Getenv("csrf"))
	req.Header.Add("priority", "u=1, i")
	req.Header.Add("sec-ch-ua", "\"Not/A)Brand\";v=\"8\", \"Chromium\";v=\"126\", \"Google Chrome\";v=\"126\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("x-li-lang", "en_US")
	req.Header.Add("x-li-pem-metadata", "Voyager - Organization - Member=organization-people-card")
	req.Header.Add("x-li-track", "{\"clientVersion\":\"1.13.19196\",\"mpVersion\":\"1.13.19196\",\"osName\":\"web\",\"timezoneOffset\":1,\"timezone\":\"Europe/London\",\"deviceFormFactor\":\"DESKTOP\",\"mpName\":\"voyager-web\",\"displayDensity\":1,\"displayWidth\":1920,\"displayHeight\":1080}")
	req.Header.Add("x-restli-protocol-version", "2.0.0")
	req.Header.Add("Cookie", os.Getenv("cookie"))
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ Error executing request:", err)
		return nil, 0
	}
	defer res.Body.Close()

	body := utils.Decoding(res)
	//fmt.Printf(string(body))
	if res.StatusCode != 200 {
		fmt.Printf("🌐 Connection Error With Status Code: %d\n", res.StatusCode)
	}
	return body, res.StatusCode
}

func GetCompanyName(client *http.Client, inputCompany string) string {
	url := "https://www.google.com/search?q=" + inputCompany + "+linkedin"
	bodyString := GetReqGoogle(url, client)

	matches := strings.Split(bodyString, "company/")[1]
	matches = strings.Split(matches, "&")[0]
	//fmt.Println(matches)
	return matches
}

func GetCompanyId(client *http.Client, companyName string) (string, error) {
	companyName = GetCompanyName(client, companyName)
	url := "https://www.linkedin.com/company/" + companyName + "/people/"

	resp, status := GetReq(url, client)
	if status != 200 {
		return "0", fmt.Errorf("error making GET request: %v", status)
	}
	// Regular expression to find the company ID
	re := regexp.MustCompile(`urn:li:fsd_company:(\d+)`)
	matches := re.FindSubmatch(resp)
	//fmt.Printf(url)
	if len(matches) < 2 {
		return "0", fmt.Errorf("company ID not found in the response")
	}

	companyID := string(matches[1])

	return companyID, nil
}

func GetReqGoogle(url string, client *http.Client) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("❌ Error creating request:", err)
		return "1 "
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ Error executing request:", err)
		return " 2"
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("🌐 Connection Errorss With Status Code: %d\n", res.StatusCode)
		fmt.Printf("url: %v\n", url)

		return "3 "
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("❌ Error reading response body:", err)
		return "4 "
	}

	return string(bodyBytes)
}

