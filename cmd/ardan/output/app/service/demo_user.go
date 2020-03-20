package service

import (
    "abc.com/def/output/app/model"
    "abc.com/def/output/app/repository"
)

type DemoUserService interface {
    Get(id string) (*model.DemoUser, error)
}
type demoUserService struct {
    repo *repository.DemoUserRepository
}

func NewDemoUserService(repo *repository.UserRepository) DemoUserService {
	return &demoUserService{repo}
}

func (s *demoUserService) Get(id string) (*model.DemoUser, error) {
	return s.repo.Get(id)
}
