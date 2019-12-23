package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/letanthang/cli_search/config"
	"github.com/letanthang/cli_search/model"
	"github.com/letanthang/cli_search/ticket"
)

type memoryTicketRepo struct {
	items []model.Ticket
}

func New() ticket.Repository {
	items := loadDataFromJSON(config.Config.Data.Ticket)
	return memoryTicketRepo{
		items: items,
	}
}

func loadDataFromJSON(filename string) []model.Ticket {
	var slice []model.Ticket
	data, err := ioutil.ReadFile(filename)
	err = json.Unmarshal(data, &slice)
	if err != nil {
		log.Fatalf("read json file error: %v", err)
	}
	//fmt.Printf("tickets: %+v", slice)
	return slice
}

func (memOrgRepo memoryTicketRepo) Describe() []string {
	return []string{"_id", "url", "external_id", "type", "subject", "description", "priority", "status", "submitter_id", "assignee_id", "tag", "has_incidents", "organization_id", "via"}
}

func (memOrgRepo memoryTicketRepo) Search(field, word string) ([]*model.Ticket, error) {
	var result []*model.Ticket
	for i, v := range memOrgRepo.items {
		ok, err := v.SearchByField(field, word)
		if err != nil {
			return result, err
		}
		if ok {
			result = append(result, &memOrgRepo.items[i])
		}
	}
	return result, nil
}
