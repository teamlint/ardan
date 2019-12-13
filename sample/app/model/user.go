package model

import (
	"time"
)

type User struct {
	Id        string `xorm:"pk"`
	Name      string
	Desc      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
