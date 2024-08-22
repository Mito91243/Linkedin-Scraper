package utils

import (
	"fmt"
	"main/internal/models"
	"net/http"
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

func GetPostQuery(people []models.ProfileRes,client *http.Client) string {
	final := ""
	for _, person := range people {
		if person.LastName == "Member" {
			continue
		}
		temp := fmt.Sprintf("%v,", strings.Split(person.Link, "%")[len(strings.Split(person.Link, "%"))-1])
		final += temp
		//fmt.Printf("Profile URN: %s, Name: %s\n", strings.Split(person.Link,"%")[len(strings.Split(person.Link,"%"))-1] , person.FullName)
	}
	url := "https://www.linkedin.com/voyager/api/graphql?variables=(start:0,origin:FACETED_SEARCH,query:(keywords:internship,flagshipSearchIntent:SEARCH_SRP,queryParameters:List((key:fromMember,value:List(" + final + ")),(key:resultType,value:List(CONTENT)),(key:sortBy,value:List(relevance))),includeFiltersInResponse:false),count:3)&queryId=voyagerSearchDashClusters.a2b606e8c1f58b3cf72fb5d54a2a57e7"
	return url
}

//https://www.linkedin.com/voyager/api/graphql?variables=(start:0,origin:FACETED_SEARCH,query:(keywords:internship,flagshipSearchIntent:SEARCH_SRP,queryParameters:List((key:fromMember,value:List(ACoAADttcpsBxvTz1m-uBKP0JaSchHJGifSGGHY,ACoAACbk0-gBoc1GMeV3vMrr8M7eUwAtK9GkvAw)),(key:resultType,value:List(CONTENT)),(key:sortBy,value:List(relevance))),includeFiltersInResponse:false),count:3)&queryId=voyagerSearchDashClusters.a2b606e8c1f58b3cf72fb5d54a2a57e7
//! Make Model for Job Post
//! Api Call with all ID's
//! Data text.text (string)// Reactions(Int) // OP name // Post link(string) // Comments (INt)
