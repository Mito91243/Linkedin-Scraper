package utils

import (
	"main/internal/models"
	"strings"
	//"fmt"
)

func SafeGetString(m map[string]interface{}, key string) string {
	if val, ok := m[key]; ok && val != nil {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

func TruncateString(s string, maxLength int) string {
	if len(s) > maxLength {
		return s[:maxLength-3] + "..."
	}
	return s
}

func GetPostsUrls(people []models.ProfileRes, companyName string, counter int) []string {
	var finalQ string
	var urls []string

	for _, profile := range people {

		//tempF := strings.Replace(profile.FullName,"."," ",3)
		tempF := strings.Replace(profile.FullName, " ", "+", 3)
		spacer := "+OR+intitle%3A%22"
		AddedName := "%22" + spacer + tempF
		finalQ = finalQ + AddedName
		counter++
		if counter == 13 {
			GetPostsUrls(people[12:], companyName, 0)
			break
		}
	}
	//fmt.Printf("Number of Valid Profiles: %v\n",counter)

	lookFor := "internship"
	url := "https://www.google.com/search?q=site%3Alinkedin.com+" + "Post" + "+intext%3A%22" + companyName + "%22+intext%3A%22" + lookFor + "%22+%28intitle%3A%22JohnDoeLOL" + finalQ + "%29"
	//fmt.Println(url)
	urls = append(urls, url)
	return urls
}
