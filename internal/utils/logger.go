package utils

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"main/internal/models"
	"os"
)

func DisplayProfiles(profiles []models.ProfileRes) {
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
	tempC := 0

	for i, profile := range profiles {
		if profile.LastName == "Member" {
			tempC++
			continue
		}
		table.Append([]string{
			fmt.Sprintf("Profile %d", i+1-tempC),
			TruncateString(profile.FullName, 20),
			TruncateString(profile.LastName, 15),
			TruncateString(profile.Position, 100),
			TruncateString(profile.ProfileURN, 40),
		})
	}
	table.Render()
}

func PrintHeader() {
	color.Set(color.FgGreen)
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println()
	fmt.Println()
	fmt.Println("  _      _        _            _  _          _____                                ")
	fmt.Println(" | |    (_)      | |          |  (_)        / ____|                               ")
	fmt.Println(" | |     _ _ __  | | _____  __| |_ _ __    | (___   ___ _ __ __ _ _ __   ___ _ __ ")
	fmt.Println(" | |    | | '_ \\| |/ / _ \\/ _` | | '_  \\   \\___ \\ / __| '__/ _` | '_ \\ / _ \\ '__|")
	fmt.Println(" | |____| | | | |   <  __/ (_| | | | | |    ____) | (__| | | (_| | |_) |  __/ |   ")
	fmt.Println(" |______|_|_| |_|_|\\_\\___|\\__,_ |_| ||_|   |_____/ \\___|_|  \\__,_| .__/ \\___|_|   ")
	fmt.Println("                                                                 | |                  ")
	fmt.Println("                                                                 |_|                  ")
	fmt.Println()
	fmt.Println()
	fmt.Println("   Welcome to LinkedIn Scraper! Enter Company Name to Start.")
	fmt.Println()
	fmt.Println("------------------------------------------------------------------------")
	color.Unset()
}
