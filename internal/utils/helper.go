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

func GetPostQuery(people []models.ProfileRes, client *http.Client) string {
	// We get the profiles we scrapped then we concatenate their Identifiers with , And remove the 3A at beginning of each ID
	final := ""
	for _, person := range people {
		if person.LastName == "Member" {
			continue
		}
		temp := fmt.Sprintf("%v,", strings.Split(person.Link, "%")[len(strings.Split(person.Link, "%"))-1])
		temp = strings.TrimPrefix(temp,"3A")
		final += temp
	}
	final = final[:len(final)-1]
	url := "https://www.linkedin.com/voyager/api/graphql?variables=(start:0,origin:FACETED_SEARCH,query:(keywords:internship,flagshipSearchIntent:SEARCH_SRP,queryParameters:List((key:fromMember,value:List("+final+")),(key:resultType,value:List(CONTENT))),includeFiltersInResponse:false),count:3)&queryId=voyagerSearchDashClusters.a2b606e8c1f58b3cf72fb5d54a2a57e7"
	fmt.Println()
	fmt.Print(url)
	fmt.Println()
	return url
}

//	url := "https://www.linkedin.com/voyager/api/graphql?variables=(start:0,origin:FACETED_SEARCH,query:(keywords:intern,flagshipSearchIntent:SEARCH_SRP,queryParameters:List((key:fromMember,value:List(ACoAADwXNCYBtzPLExyaYOEM0hHjXeHpd7MaL1U,ACoAAAmld2YBPvssug9XG9y763pbsgi7j56hQPc,ACoAACRf_ZIBibwGD4VQxL9Zx0n_iMJB57VPB6M,ACoAAD3hBFgBTRqdorxET6iP1u6GurdfW_dQ1eo,ACoAAEGJrIMBkRe5-0-OEl_Rb-NK1TpdVhXnQq0,ACoAAAynTtoBw-ue8qwFuXT3qIf0o2Z4CmoH_XA,ACoAAAAHz04BwbOUX11qxZVM2S8XYgPEWPUY1i8,ACoAABtjpPMBERe5ucxfXBJObKkgLk-IcluhdPg,ACoAAAkQF2MBYKjgfpMP1tq1rkwONr-YIonlUtE,ACoAACt4tMoBBAIj7jjoyTqHpBAkYVfu9x4Nn4E,ACoAAEPcuv0Bzmerv9dDRc_-xMlGZynWCVwAVSw,ACoAAEI_ozEBL65Ml6qPqU8LC6dYv3o6Y-mqDbA,ACoAADOKBzcB1j23WG57I9TOPtbwLPQ6Xxy5x08,ACoAACE7YVoBhFoEOgA5Z-HRT-n_zGlZObm6Y2Y)),(key:resultType,value:List(CONTENT)),(key:sortBy,value:List(relevance)))),count:10)&queryId=voyagerSearchDashClusters.a2b606e8c1f58b3cf72fb5d54a2a57e7"

//! TEsting Invalid valid
//https://www.linkedin.com/voyager/api/graphql?variables=(start:0,origin:FACETED_SEARCH,query:(keywords:internship,flagshipSearchIntent:SEARCH_SRP,queryParameters:List((key:fromMember,value:List(3AACoAAA00FZ8BEg25yRCe7vfsXx6kx59VB12ktPU,3AACoAAAz7eHwBZyzuZkl78cUp7sOSGiVDdfdbXFk)),(key:resultType,value:List(CONTENT)),(key:searchId,value:List(7d395167-cd5d-4772-8931-f56562831e0c)))),count:39)&queryId=voyagerSearchDashClusters.a2b606e8c1f58b3cf72fb5d54a2a57e7


//!ACoAACbk0-gBoc1GMeV3vMrr8M7eUwAtK9GkvAw
//