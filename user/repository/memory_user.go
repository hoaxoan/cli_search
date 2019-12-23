package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/letanthang/cli_search/config"
	"github.com/letanthang/cli_search/model"
	"github.com/letanthang/cli_search/user"
)

type memoryUserRepo struct {
	items *[]model.User
}

func New() user.Repository {
	items := loadDataFromJSON(config.Config.Data.User)
	return memoryUserRepo{
		items: &items,
	}
}

func loadDataFromJSON(filename string) []model.User {
	var slice []model.User
	data, err := ioutil.ReadFile(filename)
	err = json.Unmarshal(data, &slice)
	if err != nil {
		log.Fatalf("read json file error: %v", err)
	}
	//fmt.Printf("tickets: %+v", slice)
	return slice
}

func (memUserRepo memoryUserRepo) Describe() {

}

func (memUserRepo memoryUserRepo) Search(field, word string) ([]*model.User, error) {
	var result []*model.User
	for _, v := range *memUserRepo.items {
		if v.SearchByField(field, word) {
			result = append(result, &v)
		}
	}
	return result, nil
}
