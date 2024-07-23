package helperFunction

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"strings"
	"time"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"os"
)

func Start() {
	PrintHeader()
	err := godotenv.Load()
	if err != nil {
		fmt.Println("❌ Error loading .env file")
		return
	}

	client := &http.Client{}
	companyName := Read_input()
	companyEmail := companyName
	if len(strings.Split(companyName,"-")) > 1 {
		companyEmail = strings.Split(companyName,"-")[0]
	}
	start := time.Now()
	companyIdchan := make(chan string)
	positionIdchan := make(chan string)
	go func() {
		companyID, _ := Get_company_id(client, companyName)
		companyIdchan <- companyID
	}()
	go func() {
		positionIdentifier := ReadPositionInput()
		positionIdchan <- positionIdentifier
	}()
	companyID := <-companyIdchan
	positionIdentifier := <-positionIdchan
	// if you need a specific country to scrape from put this (key:geoUrn,value:List(102713980)) before (key:resultType in the url  . Here we getting indian people
	// The id List(id) in here is the id of the country so change it depending on what country you want to scrape from . Here is USA for reference : 103644278
	url := "https://www.linkedin.com/voyager/api/graphql?variables=(start:0,origin:FACETED_SEARCH,query:(flagshipSearchIntent:ORGANIZATIONS_PEOPLE_ALUMNI,queryParameters:List((key:currentCompany,value:List(" + companyID + ")),(key:currentFunction,value:List(" + positionIdentifier + ")),(key:geoUrn,value:List(106155005)),(key:resultType,value:List(ORGANIZATION_ALUMNI))),includeFiltersInResponse:true),count:48)&queryId=voyagerSearchDashClusters.2e313ab8de30ca45e1c025cd0cfc6199"

	profiles := Run(companyEmail,url, client)
	EncodeProfiles(profiles)
	fmt.Println(strings.Repeat("-", 60))

	fmt.Printf("✨ Total Time To Fetch Profiles: %.2f seconds\n", time.Since(start).Seconds())
	fmt.Println(strings.Repeat("=", 60))
	fmt.Scanln()
}

func Run(companyName string,url string, client *http.Client) []ProfileRes {
	start := time.Now()
	body, status := Get_Req(url, client)
	if status != 200 {
		color.Red("Error making GET request: %v", status)
		return nil
	}

	results, err := ExtractProfiles(body)
	if err != nil {
		color.Red("Error extracting profiles: %v", err)
		return nil
	}

	color.Cyan("\n🔍 Extracted Profiles:")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Profile", "Full Name", "Last Name", "Position", "Profile URL"})
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

	var profiles []ProfileRes
	for i, profile := range results {
		if _, ok := profile["position"]; !ok {
			continue
		}
		temp_profile := ProfileRes{
			FullName:   safeGetString(profile, "fullName"),
			LastName:   safeGetString(profile, "lastName"),
			Position:   safeGetString(profile, "position"),
			ProfileURN: safeGetString(profile, "Possible Email"),
		}
		emailFirst := strings.Split(temp_profile.FullName, " ")[0]
		emailLast := strings.Split(temp_profile.FullName, " ")[len(strings.Split(temp_profile.FullName, " "))-1]
		email := emailFirst + "." + emailLast + "@" + companyName + ".com"

		table.Append([]string{
			fmt.Sprintf("Profile %d", i+1-len(results)/2),
			truncateString(temp_profile.FullName, 20),
			truncateString(temp_profile.LastName, 15),
			truncateString(temp_profile.Position, 100),
			truncateString(email, 40),
		})

		profiles = append(profiles, temp_profile)
	}

	table.Render()

	color.Yellow("\n✨ Time to fetch %d profiles: %.2f seconds\n", len(profiles), time.Since(start).Seconds())
	return profiles
}