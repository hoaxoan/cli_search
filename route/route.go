package route

import (
	"encoding/json"
	"fmt"
	"github.com/letanthang/cli_search/organization"
	"github.com/letanthang/cli_search/ticket"
	"github.com/letanthang/cli_search/user"
)

var OrganizationUsecase organization.Usecase
var TicketUsecase ticket.Usecase
var UserUsecase user.Usecase

func FowardRoute(command string, args []string) {
	//validate route
	if command != "search" && command != "describe" {
		fmt.Println("wrong command [search|describe]")
		return
	}
	if command == "search" && len(args) != 3 {
		fmt.Println("wrong param: search X field value")
		return
	}
	if command == "describe" && len(args) != 1 {
		fmt.Println("wrong param: describe [user|ticket]")
		return
	}
	subject := args[0]
	Handle(command, subject, args[1:])
	// fmt.Printf("Route forward. command: %v, args: %v", command, args)
}

func Handle(command, subject string, args []string) {
	if subject == "organization" {
		if command == "describe" {
			describeOrganization(args, OrganizationUsecase)
		}
		if command == "search" {
			searchOrganization(args, OrganizationUsecase)
		}
	}

	if subject == "user" {
		if command == "describe" {
			describeUser(args, UserUsecase)
		}
		if command == "search" {
			searchUser(args, UserUsecase)
		}
	}

	if subject == "ticket" {
		if command == "describe" {
			describeTicket(args, TicketUsecase)
		}
		if command == "search" {
			searchTicket(args, TicketUsecase)
		}
	}
}

// ****** organization
func describeOrganization(args []string, uc organization.Usecase) {
	data := uc.Describe()
	fmt.Println("Organization have these field")

	for _, v := range data {
		fmt.Println(v)
	}

}

func searchOrganization(args []string, uc organization.Usecase) {
	field := args[0]
	word := args[1]
	data, err := uc.Search(field, word)
	fmt.Println("param", field, word, data)
	if err != nil {
		fmt.Printf("search organization error: %v", err)
	}

	bs, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println("Search Result")
	fmt.Println(string(bs))
}

// ****** user
func describeUser(args []string, uc user.Usecase) {

}

func searchUser(args []string, uc user.Usecase) {

}

// ****** ticket
func describeTicket(args []string, uc ticket.Usecase) {

}

func searchTicket(args []string, uc ticket.Usecase) {

}
