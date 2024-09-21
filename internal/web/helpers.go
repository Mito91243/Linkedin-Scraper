package web

import (
	"encoding/json"
	"fmt"
	"main/config"
	"main/internal/api"
	"main/internal/models"
	"main/internal/utils"
	"net/http"
	"runtime/debug"
	"strings"
)

// The serverError helper writes an error message and stack trace to the errorLog,
// then sends a generic 500 Internal Server Error response to the user.
func  serverError(app *config.Application,w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Print(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description
// to the user. We'll use this later in the book to send responses like 400 "Bad
// Request" when there's a problem with the request that the user sent.
func  clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// For consistency, we'll also implement a notFound helper. This is simply a
// convenience wrapper around clientError which sends a 404 Not Found response to
// the user.
func notFound(w http.ResponseWriter) {
	clientError(w, http.StatusNotFound)
}

func getAllProfiles(position string, companyName string, app *config.Application) []byte {

	CompanyURL := "https://www.linkedin.com/voyager/api/graphql?variables=(start:0,origin:TYPEAHEAD_ESCAPE_HATCH,query:(keywords:"+companyName+",flagshipSearchIntent:SEARCH_SRP,queryParameters:List((key:resultType,value:List(ALL))),includeFiltersInResponse:false))&queryId=voyagerSearchDashClusters.dec2e0cf0d4c89523266f6e3b44cc87c"

	body, status := api.GetReq(CompanyURL, app)
	if status != 200 {
		app.ErrorLog.Printf("Error Fetching Company URL: %v", status)
	}
	id := utils.ExtractCompanyID(body)
	//app.InfoLog.Printf("Company ID : %v ", id)

	url := "https://www.linkedin.com/voyager/api/graphql?variables=(start:0,origin:FACETED_SEARCH,query:(flagshipSearchIntent:ORGANIZATIONS_PEOPLE_ALUMNI,queryParameters:List((key:currentCompany,value:List(" + id + ")),(key:currentFunction,value:List(" + position + ")),(key:geoUrn,value:List(106155005)),(key:resultType,value:List(ORGANIZATION_ALUMNI))),includeFiltersInResponse:true),count:49)&queryId=voyagerSearchDashClusters.2e313ab8de30ca45e1c025cd0cfc6199"
	url2 := "https://www.linkedin.com/voyager/api/graphql?variables=(start:49,origin:FACETED_SEARCH,query:(flagshipSearchIntent:ORGANIZATIONS_PEOPLE_ALUMNI,queryParameters:List((key:currentCompany,value:List(" + id + ")),(key:currentFunction,value:List(" + position + ")),(key:geoUrn,value:List(106155005)),(key:resultType,value:List(ORGANIZATION_ALUMNI))),includeFiltersInResponse:true),count:49)&queryId=voyagerSearchDashClusters.2e313ab8de30ca45e1c025cd0cfc6199"
	urlTalentAcquisition := "https://www.linkedin.com/voyager/api/graphql?variables=(start:0,origin:FACETED_SEARCH,query:(keywords:Talent%20acquisition,flagshipSearchIntent:ORGANIZATIONS_PEOPLE_ALUMNI,queryParameters:List((key:currentCompany,value:List(" + id + ")),(key:geoUrn,value:List(106155005)),(key:resultType,value:List(ORGANIZATION_ALUMNI))),includeFiltersInResponse:true),count:49)&queryId=voyagerSearchDashClusters.ff737c692102a8ce842be8f129f834ae"

	firstPatch := make(chan []models.ProfileRes)
	SecondPatch := make(chan []models.ProfileRes)

	go func() {
		firstPatch <- getProfiles(companyName, url, app)
	}()
	go func() {
		SecondPatch <- getProfiles(companyName, url2, app)
	}()
	profiles := <-firstPatch
	profiles2 := <-SecondPatch
	profiles = append(profiles, profiles2...)

	// Get Talent Acquisition personnel
	if position == "12" {
		profilesExtended := getProfiles(companyName, urlTalentAcquisition, app)
		profiles = append(profiles, profilesExtended...)
	}
	//app.DB.Models.Profilesdb.InsertMany()
	jsonData, err := json.Marshal(profiles)
	if err != nil {
		fmt.Println("Error Marshalling to Json")
	}
	app.InfoLog.Printf("Profile Fetched :  %d", len(profiles))

	return jsonData
}

func  getProfiles(companyName string, url string, app *config.Application) []models.ProfileRes {

	body, status := api.GetReq(url, app)

	if status != 200 {
		app.ErrorLog.Printf("Error Getting Profiles: %v", status)
		return nil
	}

	results, err := utils.ExtractProfiles(body)
	if err != nil {
		app.ErrorLog.Printf("Error Extracting Profiles: %v", status)
		return nil
	}

	// Convert Profiles interfaces to strings and Guess emails
	var profiles []models.ProfileRes
	for _, profile := range results {
		if _, ok := profile["position"]; !ok {
			continue
		}

		temp_profile := models.ProfileRes{
			FullName:   utils.SafeGetString(profile, "fullName"),
			LastName:   utils.SafeGetString(profile, "lastName"),
			Position:   utils.SafeGetString(profile, "position"),
			ProfileURN: utils.SafeGetString(profile, "Email"),
			Link:       utils.SafeGetString(profile, "bserpEntityNavigationalUrl"),
		}

		// Predict the email of each user
		emailFirst := strings.Split(temp_profile.FullName, " ")[0]
		emailLast := strings.Split(temp_profile.FullName, " ")[len(strings.Split(temp_profile.FullName, " "))-1]
		companyName = strings.ReplaceAll(companyName, "+", "")
		email := emailFirst + "." + emailLast + "@" + companyName + ".com"
		temp_profile.ProfileURN = email
		profiles = append(profiles, temp_profile)
	}

	return profiles
}

/*func getAllPosts(profiles []models.ProfileRes, keyword string, app *config.Application) {
	posturls := []string{}

	utils.GetPostQuery(profiles, keyword, &posturls)
	if len(posturls) < 1 {
		fmt.Print("DIDNT PARSE ANY")
		return
	}
	posts := GetPosts(posturls, app)
	utils.DisplayPosts(posts)
}

func GetPosts(urls []string, app *config.Application) []models.PostRes {
	var posts []models.PostRes
	for _, url := range urls {
		body, status := api.GetReq(url, app)
		results, err := utils.ExtractPosts(body)

		if err != nil {
			fmt.Printf("Error extracting posts: %v\n", err)
			return nil
		}

		fmt.Printf("Posts Status: %d\n", status)
		posts = append(posts, results...)
	}
	return posts
}*/




//https://www.linkedin.com/voyager/api/graphql?variables=(query:paymob)&queryId=voyagerSearchDashTypeahead.d51ffbb93e101b83c05ba0734bc4f380
//https://www.linkedin.com/voyager/api/graphql?variables=(query:fawry)&queryId=voyagerSearchDashTypeahead.d51ffbb93e101b83c05ba0734bc4f380