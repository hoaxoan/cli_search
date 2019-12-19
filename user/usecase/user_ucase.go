package usecase

import (
	"github.com/letanthang/cli_search/model"
	"github.com/letanthang/cli_search/user"
)

type userUsecase struct {
	userRepo user.Repository
}

func New(o user.Repository) user.Usecase {
	return &userUsecase{
		userRepo: o,
	}
}

func (ouc *userUsecase) Describe() {

}

func (ouc *userUsecase) Search(field, word string) ([]*model.User, error) {
	return ouc.userRepo.Search(field, word)
}
