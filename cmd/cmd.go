package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"log"
	"main/internal/api"
	"main/internal/models"
	"main/internal/utils"
	"net/http"
	"strings"
	"time"
)

func Start(errorLog, infoLog *log.Logger) {
	utils.PrintHeader()
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Print(err)
		fmt.Println("❌ Error loading .env file")
		return
	}

	client := &http.Client{}
	companyName := utils.Read_input()
	start := time.Now()
	companyIdchan := make(chan string)
	positionIdchan := make(chan string)
	go func() {
		CompanyURL := "https://www.linkedin.com/voyager/api/graphql?variables=(query:" + companyName + ")&queryId=voyagerSearchDashTypeahead.5d388aa0c61a43e1dcd14aaa52fe062c"
		//fmt.Printf(" COMAPNY URL IS : %v\n",CompanyURL)
		body, status := api.GetReq(CompanyURL, client)
		if status != 200 {
			fmt.Printf("Error Fetching Company URL: %v", status)
		}
		//fmt.Print(string(body))
		id := utils.ExtractCompanyID(body)
		//fmt.Printf("ID OF THE COMPANY IS : %v ", id)
		companyIdchan <- id
	}()
	go func() {
		positionIdentifier := utils.ReadPositionInput()
		positionIdchan <- positionIdentifier
	}()
	companyID := <-companyIdchan
	positionIdentifier := <-positionIdchan
	// if you need a specific country to scrape from put this (key:geoUrn,value:List(102713980)) before (key:resultType in the url  . Here we getting indian people
	// The id List(id) in here is the id of the country so change it depending on what country you want to scrape from . Here is USA for reference : 103644278
	url := "https://www.linkedin.com/voyager/api/graphql?variables=(start:0,origin:FACETED_SEARCH,query:(flagshipSearchIntent:ORGANIZATIONS_PEOPLE_ALUMNI,queryParameters:List((key:currentCompany,value:List(" + companyID + ")),(key:currentFunction,value:List(" + positionIdentifier + ")),(key:geoUrn,value:List(106155005)),(key:resultType,value:List(ORGANIZATION_ALUMNI))),includeFiltersInResponse:true),count:49)&queryId=voyagerSearchDashClusters.2e313ab8de30ca45e1c025cd0cfc6199"
	url2 := "https://www.linkedin.com/voyager/api/graphql?variables=(start:49,origin:FACETED_SEARCH,query:(flagshipSearchIntent:ORGANIZATIONS_PEOPLE_ALUMNI,queryParameters:List((key:currentCompany,value:List(" + companyID + ")),(key:currentFunction,value:List(" + positionIdentifier + ")),(key:geoUrn,value:List(106155005)),(key:resultType,value:List(ORGANIZATION_ALUMNI))),includeFiltersInResponse:true),count:49)&queryId=voyagerSearchDashClusters.2e313ab8de30ca45e1c025cd0cfc6199"

	firstPatch := make(chan []models.ProfileRes)
	SecondPatch := make(chan []models.ProfileRes)
	go func() {
		firstPatch <- Run(companyName, url, client)
	}()
	go func() {
		SecondPatch <- Run(companyName, url2, client)
	}()
	profiles := <-firstPatch
	profiles2 := <-SecondPatch
	profiles = append(profiles, profiles2...)

	// Get Talent Acquisition personnel
	if positionIdentifier == "12" {
		urlTalentAcquisition := "https://www.linkedin.com/voyager/api/graphql?variables=(start:0,origin:FACETED_SEARCH,query:(keywords:Talent%20acquisition,flagshipSearchIntent:ORGANIZATIONS_PEOPLE_ALUMNI,queryParameters:List((key:currentCompany,value:List(" + companyID + ")),(key:geoUrn,value:List(106155005)),(key:resultType,value:List(ORGANIZATION_ALUMNI))),includeFiltersInResponse:true),count:49)&queryId=voyagerSearchDashClusters.ff737c692102a8ce842be8f129f834ae"
		profilesExtended := Run(companyName, urlTalentAcquisition, client)
		profiles = append(profiles, profilesExtended...)
	}
	utils.DisplayProfiles(profiles)
	color.Yellow("\n✨ Time to fetch %d profiles: %.2f seconds\n", len(profiles), time.Since(start).Seconds())

	fmt.Println(strings.Repeat("-", 60))

	fmt.Printf("✨ Total Time To Fetch Profiles: %.2f seconds\n", time.Since(start).Seconds())
	fmt.Println(strings.Repeat("=", 60))
	//! POSTS INCOMING STARTS HERE

	posturls := []string{}

	keyword := utils.Read_KeyWord()
	utils.GetPostQuery(profiles, keyword, &posturls)
	if len(posturls) < 1 {
		fmt.Print("DIDNT PARSE ANY")
		return
	}
	posts := GetPosts(posturls, client)
	utils.DisplayPosts(posts)
	fmt.Scanln()
}

func Run(companyName string, url string, client *http.Client) []models.ProfileRes {
	body, status := api.GetReq(url, client)
	if status != 200 {
		color.Red("Error making GET request: %v", status)
		return nil
	}

	results, err := utils.ExtractProfiles(body)
	if err != nil {
		color.Red("Error extracting profiles: %v", err)
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

func GetPosts(urls []string, client *http.Client) []models.PostRes {
	var posts []models.PostRes
	for _, url := range urls {
		body, status := api.GetReq(url, client)
		results, err := utils.ExtractPosts(body)

		if err != nil {
			fmt.Printf("Error extracting posts: %v\n", err)
			return nil
		}

		fmt.Printf("Posts Status: %d\n", status)
		posts = append(posts, results...)
	}
	return posts
}

//! Get All posts not only 24 people
