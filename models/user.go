package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
	"time"
	"github.com/asaskevich/govalidator"
)

type User struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
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