package model

import (
	"time"
)

type DemoUser struct {
	ID         string     `xorm:"not null pk unique CHAR(20) 'id'" json:"id"`
	Username   string     `xorm:"not null index VARCHAR(100)" json:"username"`
	Password   string     `xorm:"not null VARCHAR(100)" json:"password"`
	IsApproved bool       `xorm:"not null default(0)" json:"is_approved"`
	Email      string     `xorm:"not null VARCHAR(1000)" json:"email"`
	Gender     Gender     `xorm:"not null default(0) INTEGER" json:"gender"`
	Bio        string     `xorm:"not null VARCHAR(1000)" json:"bio"`
	Phone      string     `xorm:"not null VARCHAR(20)" json:"phone"`
	Posts      int64      `xorm:"not null default(0) BIGINT" json:"posts"`
	CreatedAt  time.Time  `xorm:"not null created TIMESTAMPZ" json:"created_at"`
	UpdatedAt  time.Time  `xorm:"not null updated TIMESTAMPZ" json:"updated_at"`
	DeletedAt  *time.Time `xorm:"null deleted index TIMESTAMPZ" json:"deleted_at"`
}

// TableName return table name
func (m DemoUser) TableName() string {
	return "demo_user"
}

type Gender int

const (
	GenderNotSet Gender = 0 // notset
	GenderMan    Gender = 1 // man
	GenderWoman  Gender = 2 // woman
)
