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
func (app *Application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Print(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description
// to the user. We'll use this later in the book to send responses like 400 "Bad
// Request" when there's a problem with the request that the user sent.
func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// For consistency, we'll also implement a notFound helper. This is simply a
// convenience wrapper around clientError which sends a 404 Not Found response to
// the user.
func (app *Application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *Application) getAllProfiles(position string, companyName string) []byte {
	Ap := &config.Application{
		InfoLog:  app.InfoLog,
		ErrorLog: app.ErrorLog,
		Client:   app.Client,
	}

	CompanyURL := "https://www.linkedin.com/voyager/api/graphql?variables=(start:0,origin:RICH_QUERY_TYPEAHEAD_HISTORY,query:(keywords:" + companyName + ",flagshipSearchIntent:SEARCH_SRP,queryParameters:List((key:heroEntityKey,value:List(urn%3Ali%3Aorganization%3A18057791)),(key:position,value:List(0)),(key:resultType,value:List(ALL)),(key:searchId,value:List(7f9942a4-3ddc-46ad-8d2e-e134c5d766e7)),(key:spellCorrectionEnabled,value:List(true))),includeFiltersInResponse:false,spellCorrectionEnabled:true,clientSearchId:8c0cd58c-c063-4477-9491-f25888b7987c))&queryId=voyagerSearchDashClusters.b67807cb32b49b40ee7d5f5e2310d071"

	body, status := api.GetReq(CompanyURL, Ap)
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
		firstPatch <- app.getProfiles(companyName, url, Ap)
	}()
	go func() {
		SecondPatch <- app.getProfiles(companyName, url2, Ap)
	}()
	profiles := <-firstPatch
	profiles2 := <-SecondPatch
	profiles = append(profiles, profiles2...)

	// Get Talent Acquisition personnel
	if position == "12" {
		profilesExtended := app.getProfiles(companyName, urlTalentAcquisition, Ap)
		profiles = append(profiles, profilesExtended...)
	}
	jsonData, err := json.Marshal(profiles)
	if err != nil {
		fmt.Println("Error Marshalling to Json")
	}
	app.InfoLog.Printf("Profile Fetched :  %d", len(profiles))

	return jsonData
}

func (app *Application) getProfiles(companyName string, url string, Ap *config.Application) []models.ProfileRes {

	body, status := api.GetReq(url, Ap)

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
