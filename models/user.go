package models

import (
	"gorm.io/gorm"
	"time"
	"github.com/asaskevich/govalidator"
)

type User struct {
	ID        int 		`gorm:"primary_key;autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar(255);not null" json:"username" validate:"required"`
	Email     string    `gorm:"type:varchar(255);not null" json:"email" validate:"required,email"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password" validate:"required minimum is 6 characters"`
	Photos    []Photo   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:",omitempty"`
	CreatedAt time.Time `json:",omitempty"`
	UpdatedAt time.Time `json:",omitempty"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)
	if errCreate != nil {
		err = errCreate
		return
	}
	return
}