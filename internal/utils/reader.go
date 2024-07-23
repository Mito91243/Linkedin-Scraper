package utils

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

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

func Read_input() string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("📝 Enter Company Name: ")

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			fmt.Println("❌ Empty input. Please try again.")
			fmt.Print("📝 Enter Company Name: ")

		} else {
			text = strings.Replace(text, " ", "-", 1)
			return strings.ToLower(text)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("❗ Error:", err)
	}

	fmt.Println("❌ No valid input received.")
	fmt.Println(strings.Repeat("-", 60))
	return ""
}

func ReadPositionInput() string {
	scanner := bufio.NewScanner(os.Stdin)

	color.Set(color.FgCyan)
	fmt.Println("============================================================")
	color.Set(color.FgYellow)
	fmt.Println("📝 Choose Position To Scrape:")
	fmt.Println()
	color.Set(color.FgGreen)
	fmt.Println("1️⃣  Human Resources")
	fmt.Println("2️⃣  Engineers")
	fmt.Println("3️⃣  Program Managers")
	fmt.Println("4️⃣  Research")
	fmt.Println("5️⃣  Information Technology")

	color.Set(color.FgCyan)
	fmt.Println("============================================================")
	color.Set(color.FgWhite)
	fmt.Print("Enter your choice (1-5): ")
	color.Unset()

	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "" {
			color.Set(color.FgRed)
			fmt.Println("❌ Empty input. Please try again.")
			color.Unset()
		} else {
			switch text {
			case "1":
				color.Set(color.FgYellow)
				fmt.Println("📝 Fetching Human Resources personnel...")
				color.Unset()
				return "12"
			case "2":
				color.Set(color.FgYellow)
				fmt.Println("📝 Fetching Engineers...")
				color.Unset()
				return "8"
			case "3":
				color.Set(color.FgYellow)
				fmt.Println("📝 Fetching Program Managers...")
				color.Unset()
				return "18"
			case "4":
				color.Set(color.FgYellow)
				fmt.Println("📝 Fetching Researchers...")
				color.Unset()
				return "24"
			case "5":
				color.Set(color.FgYellow)
				fmt.Println("📝 Fetching IT personnel...")
				color.Unset()
				return "13"

			default:
				color.Set(color.FgRed)
				fmt.Println("❌ Invalid input. Please enter 1, 2, or 3.")
				color.Unset()
			}
		}
		color.Set(color.FgWhite)
		fmt.Print("Enter your choice (1-5): ")
		color.Unset()
	}
	return "12"
}
