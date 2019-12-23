package usecase

import (
	"fmt"
	"github.com/letanthang/cli_search/model"
	"github.com/letanthang/cli_search/organization"
	"github.com/letanthang/cli_search/ticket"
	"github.com/letanthang/cli_search/user"
)

type ticketUsecase struct {
	ticketRepo       ticket.Repository
	userRepo         user.Repository
	organizationRepo organization.Repository
}

func New(t ticket.Repository, u user.Repository, o organization.Repository) ticket.Usecase {
	return &ticketUsecase{
		ticketRepo:       t,
		userRepo:         u,
		organizationRepo: o,
	}
}

func (uc *ticketUsecase) Describe() []string {
	return uc.ticketRepo.Describe()
}

func (uc *ticketUsecase) Search(field, word string) ([]*model.Ticket, error) {
	data, err := uc.ticketRepo.Search(field, word)
	uc.fillOrganization(data)
	uc.fillUsers(data)
	return data, err
}

func (uc *ticketUsecase) fillOrganization(data []*model.Ticket) error {
	for _, t := range data {
		orgs, err := uc.organizationRepo.Search("_id", fmt.Sprintf("%d", t.OrganizationID))
		if err != nil {
			fmt.Printf("fill organization err: %v", err)
			return err
		}
		if len(orgs) > 0 {
			t.Organization = orgs[0].Name
		}
	}
	return nil
}

func (uc *ticketUsecase) fillUsers(data []*model.Ticket) error {
	for _, t := range data {
		//fill assignee
		users, err := uc.userRepo.Search("_id", fmt.Sprintf("%d", t.AssigneeID))
		if err != nil {
			fmt.Printf("fill assignee err: %v", err)
			return err
		}
		if len(users) > 0 {
			t.Assignee = users[0].Name
		}

		//fill submitter
		users, err = uc.userRepo.Search("_id", fmt.Sprintf("%d", t.SubmitterID))
		if err != nil {
			fmt.Printf("fill submitter err: %v", err)
			return err
		}
		if len(users) > 0 {
			t.Submitter = users[0].Name
		}
	}
	return nil
}
