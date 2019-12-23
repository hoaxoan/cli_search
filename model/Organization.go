package model

import (
	"errors"
	"fmt"
	"strconv"
)

type Organization struct {
	ID            int      `json:"_id"`
	URL           string   `json:"url"`
	Name          string   `json:"name"`
	ExternalID    string   `json:"external_id"`
	DomainNames   []string `json:"domain_names"`
	Tags          []string `json:"tags"`
	SharedTickets bool     `json:"shared_tickets"`
	Details       string   `json:"details"`
	CreatedAt     string   `json:"created_at"`

	Tickets []string `json:"tickets"`
	Users   []string `json:"users"`
}

func (o Organization) SearchByField(field string, word string) (bool, error) {
	switch field {
	case "_id":
		return o.SearchID(word), nil
	case "url":
		return o.SearchURL(word), nil
	case "name":
		return o.SearchName(word), nil
	case "external_id":
		return o.SearchExternalID(word), nil
	case "tag":
		return o.SearchTag(word), nil
	case "domain_name":
		return o.SearchDomainNames(word), nil
	case "shared_ticket":
		return o.SearchSharedTicket(word), nil
	default:
		fmt.Println("Unsupported field search:", field)

		return false, errors.New("Unsupported field search:" + field)
	}
}

func (o Organization) SearchID(word string) bool {
	ID, err := strconv.Atoi(word)
	if err != nil {
		fmt.Println(err)
	}
	if o.ID == ID {
		return true
	}
	return false
}

func (o Organization) SearchURL(word string) bool {
	if o.URL == word {
		return true
	}
	return false
}

func (o Organization) SearchName(word string) bool {
	if o.Name == word {
		return true
	}
	return false
}

func (o Organization) SearchExternalID(word string) bool {
	if o.ExternalID == word {
		return true
	}
	return false
}

func (o Organization) SearchSharedTicket(word string) bool {
	sharedTicket, err := strconv.ParseBool(word)
	if err != nil {
		fmt.Println(err)
	}
	if o.SharedTickets == sharedTicket {
		return true
	}
	return false
}

func (o Organization) SearchTag(word string) bool {
	for _, v := range o.Tags {
		if v == word {
			return true
		}
	}
	return false
}

func (o Organization) SearchDomainNames(word string) bool {
	for _, v := range o.DomainNames {
		if v == word {
			return true
		}
	}
	return false
}
