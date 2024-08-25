package utils

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"main/internal/models"
	"os"
	"strings"
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

func DisplayPosts(posts []models.PostRes) {
	color.Cyan("\n📝 Extracted Posts:")

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Post", "Author", "Content", "Reactions", "Comments", "URN", "Date", "Action Target"})
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.FgHiBlueColor},
		tablewriter.Colors{tablewriter.FgBlueColor},
		tablewriter.Colors{tablewriter.FgYellowColor},
		tablewriter.Colors{tablewriter.FgMagentaColor},
		tablewriter.Colors{tablewriter.FgCyanColor},
		tablewriter.Colors{tablewriter.FgHiMagentaColor},
		tablewriter.Colors{tablewriter.FgHiYellowColor},
	)

	table.SetAutoWrapText(false)
	table.SetColumnAlignment([]int{
		tablewriter.ALIGN_LEFT, 
		tablewriter.ALIGN_LEFT, 
		tablewriter.ALIGN_LEFT, 
		tablewriter.ALIGN_RIGHT, 
		tablewriter.ALIGN_RIGHT, 
		tablewriter.ALIGN_LEFT, 
		tablewriter.ALIGN_LEFT,
		tablewriter.ALIGN_LEFT,
	})
	table.SetColWidth(150) // Adjust this value as needed

	for i, post := range posts {
		urn := TruncateString(strings.TrimPrefix(post.URN, "urn:li:fsd_update:(urn:li:activity:"), 20)
		content := strings.ReplaceAll(post.Text, "\n", " ")
		date := TruncateString(post.Date, 30) // Truncate date if it's too long
		if date == "" {
			date = "N/A"
		}
		actionTarget := TruncateString(post.ActionTarget, 30) // Truncate ActionTarget if it's too long
		if actionTarget == "" {
			actionTarget = "N/A"
		}
		table.Append([]string{
			fmt.Sprintf("Post %d", i+1),
			TruncateString(post.Name, 20),
			TruncateString(content, 100),
			fmt.Sprintf("%d", post.NumLikes),
			fmt.Sprintf("%d", post.NumComments),
			urn,
			date,
			actionTarget,
		})
	}

	table.Render()

	fmt.Printf("\nTotal posts extracted: %d\n", len(posts))
}

