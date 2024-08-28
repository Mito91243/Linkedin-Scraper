package utils

import (
	"fmt"
	"main/internal/models"
	//"net/http"
	"strings"
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

func GetPostQuery(people []models.ProfileRes, key string, urls *[]string) {
    final := ""
    count := 0

    for _, person := range people {
        if person.LastName == "Member" {
            continue
        }

        temp := strings.TrimPrefix(strings.Split(person.Link, "%")[len(strings.Split(person.Link, "%"))-1], "3A")
        final += temp + ","

        count++
        if count == 25 {
            break
        }
    }

    if final != "" {
        final = final[:len(final)-1] // Remove the trailing comma
        url := fmt.Sprintf("https://www.linkedin.com/voyager/api/graphql?variables=(start:0,origin:FACETED_SEARCH,query:(keywords:%s,flagshipSearchIntent:SEARCH_SRP,queryParameters:List((key:fromMember,value:List(%s)),(key:resultType,value:List(CONTENT))),includeFiltersInResponse:false),count:48)&queryId=voyagerSearchDashClusters.a2b606e8c1f58b3cf72fb5d54a2a57e7", key, final)
        *urls = append(*urls, url)
    }

    if len(people) > 25 {
        GetPostQuery(people[25:], key, urls)
    }
}