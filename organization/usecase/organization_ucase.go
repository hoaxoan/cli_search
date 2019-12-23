package usecase

import (
	"fmt"
	"github.com/letanthang/cli_search/model"
	"github.com/letanthang/cli_search/organization"
	"github.com/letanthang/cli_search/ticket"
	"github.com/letanthang/cli_search/user"
)

type organizationUsecase struct {
	organizationRepo organization.Repository
	userRepo         user.Repository
	ticketRepo       ticket.Repository
}

func New(o organization.Repository, t ticket.Repository, u user.Repository) organization.Usecase {
	return &organizationUsecase{
		organizationRepo: o,
		userRepo:         u,
		ticketRepo:       t,
	}
}

func (ouc *organizationUsecase) Describe() []string {
	return ouc.organizationRepo.Describe()
}

func (ouc *organizationUsecase) Search(field, word string) ([]*model.Organization, error) {
	data, err := ouc.organizationRepo.Search(field, word)
	err = ouc.fillTicket(data)
	err = ouc.fillUser(data)
	return data, err
}

func (ouc *organizationUsecase) fillUser(data []*model.Organization) error {

	for _, o := range data {
		users, err := ouc.userRepo.Search("organization_id", fmt.Sprintf("%d", o.ID))
		if err != nil {
			fmt.Printf("fill user err %v", err)
			return err
		}

		for _, u := range users {
			o.Users = append(o.Users, u.Name)
		}
	}
	return nil
}

func (ouc *organizationUsecase) fillTicket(data []*model.Organization) error {
	for _, o := range data {
		tickets, err := ouc.ticketRepo.Search("organization_id", fmt.Sprintf("%d", o.ID))
		if err != nil {
			fmt.Printf("fill ticket err %v", err)
			return err
		}
		for _, t := range tickets {
			o.Tickets = append(o.Tickets, t.Subject)
		}
	}
	return nil
}
