package usecase

import (
	"fmt"

	"github.com/letanthang/cli_search/model"
	"github.com/letanthang/cli_search/organization"
	"github.com/letanthang/cli_search/ticket"
	"github.com/letanthang/cli_search/user"
)

type userUsecase struct {
	userRepo         user.Repository
	ticketRepo       ticket.Repository
	organizationRepo organization.Repository
}

func New(u user.Repository, t ticket.Repository, o organization.Repository) user.Usecase {
	return &userUsecase{
		userRepo:         u,
		ticketRepo:       t,
		organizationRepo: o,
	}
}

func (uc *userUsecase) Describe() []string {
	return uc.userRepo.Describe()
}

func (uc *userUsecase) Search(field, word string) ([]*model.User, error) {
	var data []*model.User
	data, err := uc.userRepo.Search(field, word)
	if err != nil {
		return data, err
	}
	uc.fillOrganization(data)
	uc.fillTicket(data)
	return data, nil
}

func (uc *userUsecase) fillOrganization(data []*model.User) error {
	for _, u := range data {
		orgs, err := uc.organizationRepo.Search("_id", fmt.Sprintf("%d", u.OrganizationID))
		if err != nil {
			fmt.Printf("fill organization err: %v", err)
			return err
		}
		if len(orgs) > 0 {
			u.Organization = orgs[0].Name
		}
	}
	return nil
}

func (uc *userUsecase) fillTicket(data []*model.User) error {
	for _, u := range data {
		//fill assignee ticket
		tickets, err := uc.ticketRepo.Search("assignee_id", fmt.Sprintf("%d", u.ID))
		if err != nil {
			fmt.Printf("fill assignee err: %v", err)
			return err
		}

		for _, t := range tickets {
			u.AssigneeTickets = append(u.AssigneeTickets, t.Subject)
		}

		//fill submitter ticket
		tickets, err = uc.ticketRepo.Search("submitter_id", fmt.Sprintf("%d", u.ID))
		if err != nil {
			fmt.Printf("fill assignee err: %v", err)
			return err
		}

		for _, t := range tickets {
			u.AssigneeTickets = append(u.SubmitterTickets, t.Subject)
		}
	}
	return nil
}
