package repository

import (
	"errors"
	"log"
	"abc.com/def/output/app/model"
	"xorm.io/xorm"
)

type DemoUserRepository struct {
	db *xorm.Engine
}

func NewDemoUserRepository(db *xorm.Engine) *DemoUserRepository {
	return &DemoUserRepository{db}
}
func (r *DemoUserRepository) Get(id string) (*model.DemoUser, error) {
	var item model.DemoUser
	has, err := r.db.ID(id).Get(&item)
	log.Printf("[DemoUserRepository.Get] has: %v, item: %+v, err: %+v", has, item, err)
	if has {
		return &item, err
	}
	return nil, errors.New("not found")
}
