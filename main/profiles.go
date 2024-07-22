package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"helperFunctions/helper"
	"net/http"
	"strings"
	"time"
)

func main() {
	helperFunction.PrintHeader()
	err := godotenv.Load()
	if err != nil {
		fmt.Println("❌ Error loading .env file")
		return
	}

	client := &http.Client{}
	companyName := helperFunction.Read_input()
	start := time.Now()
	companyIdchan := make(chan string)
	positionIdchan := make(chan string)
	go func() {
		companyID, _ := helperFunction.Get_company_id(client, companyName)
		companyIdchan <- companyID
	}()
	go func() {
		positionIdentifier := helperFunction.ReadPositionInput()
		positionIdchan <- positionIdentifier
	}()
	companyID := <-companyIdchan
	positionIdentifier := <-positionIdchan
	// if you need a specific country to scrape from put this (key:geoUrn,value:List(102713980)) before (key:resultType in the url  . Here we getting indian people
	// The id List(id) in here is the id of the country so change it depending on what country you want to scrape from . Here is USA for reference : 103644278
	url := "https://www.linkedin.com/voyager/api/graphql?variables=(start:0,origin:FACETED_SEARCH,query:(flagshipSearchIntent:ORGANIZATIONS_PEOPLE_ALUMNI,queryParameters:List((key:currentCompany,value:List(" + companyID + ")),(key:currentFunction,value:List(" + positionIdentifier + ")),(key:resultType,value:List(ORGANIZATION_ALUMNI))),includeFiltersInResponse:true),count:48)&queryId=voyagerSearchDashClusters.2e313ab8de30ca45e1c025cd0cfc6199"

	profiles := helperFunction.Run(url, client)
	helperFunction.EncodeProfiles(profiles)
	fmt.Println(strings.Repeat("-", 60))

	fmt.Printf("✨ Total Time To Fetch Profiles: %.2f seconds\n", time.Since(start).Seconds())
	fmt.Println(strings.Repeat("=", 60))
	fmt.Scanln()
}
