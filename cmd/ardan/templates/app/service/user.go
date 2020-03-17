package service

import (
	"github.com/teamlint/ardan/sample/app/model"
	"github.com/teamlint/ardan/sample/app/repository"
)

type UserService interface {
	Get(id string) (*model.User, error)
}
type userService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) UserService {
	return &userService{repo}
}

func (s *userService) Get(id string) (*model.User, error) {
	return s.repo.Get(id)
}
