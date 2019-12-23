package model

import (
	"errors"
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

	Organization string `json:"organization"`
	Submitter    string `json:"submitter"`
	Assignee     string `json:"assignee"`
}

func (t Ticket) SearchByField(field string, word string) (bool, error) {
	switch field {
	case "_id":
		return t.SearchID(word), nil
	case "url":
		return t.SearchURL(word), nil
	case "subject":
		return t.SearchSubject(word), nil
	case "external_id":
		return t.SearchExternalID(word), nil
	case "tag":
		return t.SearchTag(word), nil

	case "organization_id":
		return t.SearchOrganizationID(word), nil
	case "assignee_id":
		return t.SearchAssigneeID(word), nil
	case "submitter_id":
		return t.SearchSubmitterID(word), nil
	default:
		fmt.Println("Unsupported field search:", field)
		return false, errors.New("Unsupported field search: " + field)
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

func (t Ticket) SearchSubmitterID(word string) bool {
	submitterID, err := strconv.Atoi(word)
	if err != nil {
		fmt.Println(err)
	}
	if t.SubmitterID == submitterID {
		return true
	}
	return false
}

func (t Ticket) SearchAssigneeID(word string) bool {
	assigneeID, err := strconv.Atoi(word)
	if err != nil {
		fmt.Println(err)
	}
	if t.AssigneeID == assigneeID {
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
