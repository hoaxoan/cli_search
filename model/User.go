package model

import (
	"fmt"
	"strconv"
)

type User struct {
	ID             int      `json:"_id"`
	URL            string   `json:"url"`
	ExternalID     string   `json:"external_id"`
	Name           string   `json:"name"`
	Alias          string   `json:"alias"`
	Tags           []string `json:"tags"`
	OrganizationID int      `json:"organization_id"`
	Role           string   `json:"role"`
	Shared         bool     `json:"shared"`
	Locale         string   `json:"locale"`
	Verified       bool     `json:"verified"`
	Active         bool     `json:"active"`
	Suspended      bool     `json:"suspended"`
	CreatedAt      string   `json:"created_at"`

	Organization  string   `json:"organization"`
	AssignTickets []string `json:"assign_tickets"`
}

func (u User) SearchByField(field string, word string) bool {
	switch field {
	case "_id":
		return u.SearchID(word)
	case "url":
		return u.SearchURL(word)
	case "subject":
		return u.SearchName(word)
	case "external_id":
		return u.SearchExternalID(word)
	case "tag":
		return u.SearchTag(word)

	case "organization_id":
		return u.SearchOrganizationID(word)
	default:
		fmt.Println("Unsupported field search:", field)
		return false
	}
}

func (u User) SearchID(word string) bool {
	ID, err := strconv.Atoi(word)
	if err != nil {
		fmt.Println(err)
	}
	if u.ID == ID {
		return true
	}
	return false
}

func (u User) SearchURL(word string) bool {
	if u.URL == word {
		return true
	}
	return false
}

func (u User) SearchName(word string) bool {
	if u.Name == word {
		return true
	}
	return false
}

func (u User) SearchExternalID(word string) bool {
	if u.ExternalID == word {
		return true
	}
	return false
}

func (u User) SearchOrganizationID(word string) bool {
	organizationID, err := strconv.Atoi(word)
	if err != nil {
		fmt.Println(err)
	}
	if u.OrganizationID == organizationID {
		return true
	}
	return false
}

func (u User) SearchTag(word string) bool {
	for _, v := range u.Tags {
		if v == word {
			return true
		}
	}
	return false
}
