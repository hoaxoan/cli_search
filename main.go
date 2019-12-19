package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/letanthang/cli_search/model"
)

func check(err error) {
	if err != nil {
		log.Fatalf("App failed at err: %v", err)
	}
}

var (
	organizations []model.Organization
	tickets       []model.Ticket
	users         []model.User

	orgs []map[string]interface{}
)

func main() {
	// loadOrganization()
	loadTicket()
	// loadUser()

	loadOrgToMap()

}

func loadOrganization() {
	fp := "data/organizations.json"
	data, err := ioutil.ReadFile(fp)
	check(err)
	err = json.Unmarshal(data, &organizations)
	check(err)
	fmt.Printf("organizations: %+v", organizations)
}
func loadOrgToMap() {
	fp := "data/organizations.json"
	data, err := ioutil.ReadFile(fp)
	check(err)
	err = json.Unmarshal(data, &orgs)
	check(err)
	fmt.Printf("orgs[0]: %+v", int(orgs[0]["_id"].(float64)))

}

func searchOrganization(field, word string) {

}

func loadTicket() {
	fp := "data/tickets.json"
	data, err := ioutil.ReadFile(fp)
	check(err)
	err = json.Unmarshal(data, &tickets)
	check(err)
	fmt.Printf("tickets: %+v", tickets)
}
func loadUser() {
	fp := "data/users.json"
	data, err := ioutil.ReadFile(fp)
	check(err)
	err = json.Unmarshal(data, &users)
	check(err)
	fmt.Printf("users: %+v", users)
}
