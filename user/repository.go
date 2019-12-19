package user

import (
	"github.com/letanthang/cli_search/model"
)

type Repository interface {
	Describe()
	Search(field, word string) ([]*model.User, error)
}