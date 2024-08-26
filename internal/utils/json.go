package utils

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"main/internal/models"
	"net/http"
	"regexp"
	"strings"
)

func Decoding(res *http.Response) []byte {

	var body []byte
	var err error
	defer res.Body.Close()

	if res.Header.Get("Content-Encoding") == "gzip" {
		gzipReader, err := gzip.NewReader(res.Body)
		if err != nil {
			fmt.Println("Failed to create gzip reader:", err)
			return nil
		}
		defer gzipReader.Close()
		body, err = io.ReadAll(gzipReader)
		if err != nil {
			fmt.Println("Failed to read gzipped body:", err)
			return nil
		}
	} else {
		body, err = io.ReadAll(res.Body)
		if err != nil {
			fmt.Println("Failed to read response body:", err)
			return nil
		}
	}
	return body
}

// For Future API
func EncodeProfiles(profiles []models.ProfileRes) {
	_, err := json.Marshal(profiles)
	if err != nil {
		return
	}
}

func ExtractProfiles(jsonData []byte) ([]map[string]interface{}, error) {
	var resp models.AutoGenerated
	err := json.Unmarshal(jsonData, &resp)
	if err != nil {
		return nil, err
	}

	var profiles []map[string]interface{}

	for _, item := range resp.Included {
		if strings.Contains(item.EntityUrn, "urn:li:fsd_profile:") {
			profile := make(map[string]interface{})
			profile["entityUrn"] = item.NavigationURL
			fullName := item.Title.Text
			profile["fullName"] = fullName
			profile["bserpEntityNavigationalUrl"] = item.BserpEntityNavigationalURL
			nameParts := strings.Split(fullName, " ")
			if len(nameParts) > 0 {
				profile["lastName"] = nameParts[len(nameParts)-1]
			}

			if item.PrimarySubtitle.Text != "" {
				profile["position"] = item.PrimarySubtitle.Text
			}

			profiles = append(profiles, profile)
		}
	}

	return profiles, nil
}

func ExtractPosts(jsonData []byte) ([]models.PostRes, error) {
	var resp models.Post
	err := json.Unmarshal(jsonData, &resp)
	if err != nil {
		return nil, err
	}

	var posts []models.PostRes
	socialCounts := make(map[string]models.SocialCounts)

	// Regular expression to find https:// links
	linkRegex := regexp.MustCompile(`https://[^\s]+`)

	// First, extract all social activity counts
	for _, item := range resp.Included {
		if strings.Contains(item.EntityUrn, "socialActivityCounts") {
			urn := item.Urn
			id := extractID(urn)
			socialCounts[id] = models.SocialCounts{
				NumLikes:    item.NumLikes,
				NumComments: item.NumComments,
			}
		}
	}

	// Then, extract all posts and match with social counts
	for _, item := range resp.Included {
		if strings.Contains(item.EntityUrn, "fsd_update") {
			post := models.PostRes{
				URN:  item.EntityUrn,
				Name: item.Actor.Name.Text,
				Text: item.Commentary.Text.Text,
				Date: item.Actor.SubDescription.AccessibilityText,
			}

			if post.Text == "" && item.Content.ArticleComponent.Title.Text != "" {
				post.Text = item.Content.ArticleComponent.Title.Text
			}

			// Extract https:// link from the text
			links := linkRegex.FindAllString(post.Text, -1)
			if len(links) > 0 {
				post.ActionTarget = links[0] // Store the first link found
			}

			// If no link found in the text, use the existing ActionTarget
			if post.ActionTarget == "" && item.Content.ArticleComponent.NavigationContext.ActionTarget != "" {
				post.ActionTarget = item.Content.ArticleComponent.NavigationContext.ActionTarget
			}

			// Match social counts using the shareUrn from metadata
			if item.Metadata.ShareUrn != "" {
				shareID := extractID(item.Metadata.ShareUrn)
				if counts, ok := socialCounts[shareID]; ok {
					post.NumLikes = counts.NumLikes
					post.NumComments = counts.NumComments
				} else {
					// Try matching with the activity ID from the entityUrn
					activityID := extractActivityID(item.EntityUrn)
					if counts, ok := socialCounts[activityID]; ok {
						post.NumLikes = counts.NumLikes
						post.NumComments = counts.NumComments
					}
				}
			}

			if post.Name != "" && post.Text != "" {
				posts = append(posts, post)
			}
		}
	}

	return posts, nil
}

func extractID(urn string) string {
	parts := strings.Split(urn, ":")
	if len(parts) > 0 {
		return parts[len(parts)-1]
	}
	return ""
}

func extractActivityID(urn string) string {
	start := strings.Index(urn, "activity:")
	if start != -1 {
		start += 9 // length of "activity:"
		end := strings.Index(urn[start:], ",")
		if end != -1 {
			return urn[start : start+end]
		}
		return urn[start:]
	}
	return ""
}
