package model

import (
	"fmt"
	"strconv"
)

type Ticket struct {
	ID             string   `json:"_id"`
	URL            string   `json:"url"`
	ExternalID     string   `json:"external_id"`
	Type           string   `json:"type"`
	Subject        string   `json:"subject"`
	Description    string   `json:"description"`
	Priority       string   `json:"priority"`
	Status         string   `json:"status"`
	SubmitterID    int      `json:"submitter_id"`
	AssigneeID     int      `json:"assignee_id"`
	Tags           []string `json:"tags"`
	HasIncidents   bool     `json:"has_incidents"`
	OrganizationID int      `json:"organization_id"`
	CreatedAt      string   `json:"created_at"`
	DueAt          string   `json:"due_at"`
	Via            string   `json:"via"`

	Submitter string `json:"submitter"`
	Assignee  string `json:"assignee"`
	Ticket    string `json:"organization"`
}

func (t Ticket) SearchByField(field string, word string) bool {
	switch field {
	case "_id":
		return t.SearchID(word)
	case "url":
		return t.SearchURL(word)
	case "subject":
		return t.SearchSubject(word)
	case "external_id":
		return t.SearchExternalID(word)
	case "tag":
		return t.SearchTag(word)

	case "organization_id":
		return t.SearchOrganizationID(word)
	default:
		fmt.Println("Unsupported field search:", field)
		return false
	}
}

func (t Ticket) SearchID(word string) bool {
	if t.ID == word {
		return true
	}
	return false
}

func (t Ticket) SearchURL(word string) bool {
	if t.URL == word {
		return true
	}
	return false
}

func (t Ticket) SearchSubject(word string) bool {
	if t.Subject == word {
		return true
	}
	return false
}

func (t Ticket) SearchExternalID(word string) bool {
	if t.ExternalID == word {
		return true
	}
	return false
}

func (t Ticket) SearchOrganizationID(word string) bool {
	organizationID, err := strconv.Atoi(word)
	if err != nil {
		fmt.Println(err)
	}
	if t.OrganizationID == organizationID {
		return true
	}
	return false
}

func (t Ticket) SearchTag(word string) bool {
	for _, v := range t.Tags {
		if v == word {
			return true
		}
	}
	return false
}
