package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// LinkedIn login endpoint
	loginURL := "https://www.linkedin.com/checkpoint/lg/login-submit"

	// Create form data
	data := url.Values{}
	data.Set("session_key", "pitoamir6@gmail.com")
	data.Set("session_password", "Clara12345678")
	// Add other form fields as needed

	// Create HTTP client
	client := &http.Client{}

	// Create request
	req, err := http.NewRequest("POST", loginURL, strings.NewReader(data.Encode()))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36")

	// Send request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read response body (optional, for debugging)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	// Extract CSRF token from response headers
	csrfToken := resp.Header.Get("X-CSRF-Token")
	fmt.Println("CSRF Token:", csrfToken)

	// Extract and combine all cookies
	var cookieStrings []string
	for _, cookie := range resp.Cookies() {
		cookieStrings = append(cookieStrings, cookie.Name+"="+cookie.Value)
	}
	fullCookie := strings.Join(cookieStrings, "; ")
	fmt.Println("Full Cookie:", fullCookie)

	// Print response body (optional, for debugging)
	fmt.Println("Response Body:", string(body))
}
