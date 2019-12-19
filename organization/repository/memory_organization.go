package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/letanthang/cli_search/config"
	"github.com/letanthang/cli_search/model"
	"github.com/letanthang/cli_search/organization"
)

type memoryOrganizationRepo struct {
	items *[]model.Organization
}

func New() organization.Repository {
	items := loadDataFromJSON(config.Config.Data.Organization)
	return &memoryOrganizationRepo{
		items: &items,
	}
}

func loadDataFromJSON(filename string) []model.Organization {
	var slice []model.Organization
	data, err := ioutil.ReadFile(filename)
	err = json.Unmarshal(data, &slice)
	if err != nil {
		log.Fatalf("read json file error: %v", err)
	}
	fmt.Printf("organizations: %+v", slice)
	return slice
}

func (memOrgRepo *memoryOrganizationRepo) Describe() {

}

func (memOrgRepo *memoryOrganizationRepo) Search(field, word string) ([]*model.Organization, error) {
	var result []*model.Organization
	for _, v := range *memOrgRepo.items {
		if v.SearchByField(field, word) {
			result = append(result, &v)
		}
	}
	return result, nil
}
