package server

import (
	"fmt"
	"main/internal/api"
	"main/internal/models"
	"main/internal/utils"
	"net/http"
	"os"
	"strings"
	//"sync"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/olekukonko/tablewriter"
)

func Start() {
	utils.PrintHeader()
	err := godotenv.Load("./.env")
	if err != nil {
		fmt.Print(err)
		fmt.Println("❌ Error loading .env file")
		return
	}

	client := &http.Client{}
	companyName := utils.Read_input()
	companyEmail := companyName
	if len(strings.Split(companyName, "-")) > 1 {
		companyEmail = strings.Split(companyName, "-")[0]
	}
	start := time.Now()
	companyIdchan := make(chan string)
	positionIdchan := make(chan string)
	go func() {
		companyID, _ := api.GetCompanyId(client, companyName)
		companyIdchan <- companyID
	}()
	go func() {
		positionIdentifier := utils.ReadPositionInput()
		positionIdchan <- positionIdentifier
	}()
	companyID := <-companyIdchan
	positionIdentifier := <-positionIdchan
	// if you need a specific country to scrape from put this (key:geoUrn,value:List(102713980)) before (key:resultType in the url  . Here we getting indian people
	// The id List(id) in here is the id of the country so change it depending on what country you want to scrape from . Here is USA for reference : 103644278
	url := "https://www.linkedin.com/voyager/api/graphql?variables=(start:0,origin:FACETED_SEARCH,query:(flagshipSearchIntent:ORGANIZATIONS_PEOPLE_ALUMNI,queryParameters:List((key:currentCompany,value:List(" + companyID + ")),(key:currentFunction,value:List(" + positionIdentifier + ")),(key:geoUrn,value:List(106155005)),(key:resultType,value:List(ORGANIZATION_ALUMNI))),includeFiltersInResponse:true),count:48)&queryId=voyagerSearchDashClusters.2e313ab8de30ca45e1c025cd0cfc6199"

	profiles := Run(companyEmail, url, client)

	//utils.EncodeProfiles(profiles)
	fmt.Println(strings.Repeat("-", 60))

	fmt.Printf("✨ Total Time To Fetch Profiles: %.2f seconds\n", time.Since(start).Seconds())
	fmt.Println(strings.Repeat("=", 60))
	//! Call a generic Get_Req_Google With GO and pass a chan string to post or link in post if possible
	//! After Trial & Error This is not possible due google rate limiting
	//? Maybe Make 2 URL Calls and sleep for 30-40 Sec ?
	urls := utils.GetPostsUrls(profiles, companyName, 0)
	for _ , url := range urls {
		fmt.Println(url)
	}

	//!
	fmt.Scanln()
}

func Run(companyName string, url string, client *http.Client) []models.ProfileRes {
	start := time.Now()
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

	color.Cyan("\n🔍 Extracted Profiles:")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Profile", "Full Name", "Last Name", "Position", "Email"})
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.FgHiBlueColor},
		tablewriter.Colors{tablewriter.FgBlueColor},
		tablewriter.Colors{tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.FgMagentaColor},
	)

	table.SetAutoWrapText(false)
	table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT, tablewriter.ALIGN_LEFT})
	table.SetColWidth(50)

	var profiles []models.ProfileRes
	tempC := 0
	for i, profile := range results {
		if _, ok := profile["position"]; !ok {
			continue
		}

		temp_profile := models.ProfileRes{
			FullName:   utils.SafeGetString(profile, "fullName"),
			LastName:   utils.SafeGetString(profile, "lastName"),
			Position:   utils.SafeGetString(profile, "position"),
			ProfileURN: utils.SafeGetString(profile, "Email"),
		}
		if temp_profile.LastName == "Member" {
			tempC++
			continue
		}
		emailFirst := strings.Split(temp_profile.FullName, " ")[0]
		emailLast := strings.Split(temp_profile.FullName, " ")[len(strings.Split(temp_profile.FullName, " "))-1]
		email := emailFirst + "." + emailLast + "@" + companyName + ".com"

		table.Append([]string{
			fmt.Sprintf("Profile %d", i-tempC+1-len(results)/2),
			utils.TruncateString(temp_profile.FullName, 20),
			utils.TruncateString(temp_profile.LastName, 15),
			utils.TruncateString(temp_profile.Position, 100),
			utils.TruncateString(email, 40),
		})

		profiles = append(profiles, temp_profile)
	}

	table.Render()

	color.Yellow("\n✨ Time to fetch %d profiles: %.2f seconds\n", len(profiles), time.Since(start).Seconds())
	return profiles
}
