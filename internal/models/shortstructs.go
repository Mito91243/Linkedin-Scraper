package models

import "time"

type ProfileRes struct {
    FullName          string    `json:"fullName"`
    LastName          string    `json:"lastName"`
    Position          string    `json:"position"`
    ProfileURN        string    `json:"Email"`
    Link              string    `json:"Link"`
    ID                int64     `json:"-"`
    CreatedAt         time.Time `json:"-"`
    Version           int       `json:"-"`
    CompanyID         int       `json:"companyID"`
    Category          int       `json:"category"`
}
type PostRes struct {
	Text         string `json:"text"`
	Name         string `json:"name"`
	ActionTarget string `json:"actionTarget"`
	URN          string `json:"urn"`
	NumLikes     int    `json:"numLikes"`
	NumComments  int    `json:"numComments"`
	Date         string `json:"date"` // Changed to string
}

type SocialCounts struct {
	NumLikes    int
	NumComments int
}

type Response struct {
	Data struct {
		Included []struct {
			EntityUrn string `json:"entityUrn"`
		} `json:"included"`
	} `json:"data"`
}
