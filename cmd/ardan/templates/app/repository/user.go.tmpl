package repository

import (
	"errors"
	"log"

	"github.com/teamlint/ardan/sample/app/model"
	"xorm.io/xorm"
)

type UserRepository struct {
	db *xorm.Engine
}

func NewUserRepository(db *xorm.Engine) *UserRepository {
	return &UserRepository{db}
}
func (r *UserRepository) Get(id string) (*model.User, error) {
	var item model.User
	has, err := r.db.ID(id).Get(&item)
	log.Printf("[UserRepository.Get] has: %v, item: %+v, err: %+v", has, item, err)
	if has {
		return &item, err
	}
	return nil, errors.New("not found")
}
