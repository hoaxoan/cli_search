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

func (memOrgRepo memoryTicketRepo) Describe() {

}

func (memOrgRepo memoryTicketRepo) Search(field, word string) ([]*model.Ticket, error) {
	var result []*model.Ticket
	for i, v := range memOrgRepo.items {
		if v.SearchByField(field, word) {
			result = append(result, &memOrgRepo.items[i])
		}
	}
	return result, nil
}
