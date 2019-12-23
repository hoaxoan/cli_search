package ticket

import (
	"github.com/letanthang/cli_search/model"
)

type Repository interface {
	Describe() []string
	Search(field, word string) ([]*model.Ticket, error)
}
