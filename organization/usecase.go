package organization

import (
	"github.com/letanthang/cli_search/model"
)

type Usecase interface {
	Describe() []string
	Search(field, word string) ([]*model.Organization, error)
}
