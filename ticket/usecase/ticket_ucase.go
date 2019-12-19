package usecase

import (
	"github.com/letanthang/cli_search/model"
	"github.com/letanthang/cli_search/ticket"
)

type ticketUsecase struct {
	ticketRepo ticket.Repository
}

func New(o ticket.Repository) ticket.Usecase {
	return &ticketUsecase{
		ticketRepo: o,
	}
}

func (ouc *ticketUsecase) Describe() {

}

func (ouc *ticketUsecase) Search(field, word string) ([]*model.Ticket, error) {
	return ouc.ticketRepo.Search(field, word)
}
