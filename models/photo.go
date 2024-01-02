package models

import (
	"gorm.io/gorm"
	"time"
	"github.com/asaskevich/govalidator"	
)

type Photo struct {
	ID        int 		`gorm:"primary_key;" json:"id"`
	Title     string    `gorm:"type:varchar(255);not null" json:"title" validate:"required"`
	Caption   string    `gorm:"type:varchar(255);not null" json:"caption" validate:"required"`
	PhotoUrl  string    `gorm:"type:varchar(255);not null" json:"photo_url" validate:"required"`
	UserID    int		`json:"user_id"`
	User      *User     `json:"user"`
	CreatedAt time.Time `json:",omitempty"`
	UpdatedAt time.Time `json:",omitempty"`
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}
	return
}

func (p *Photo) BeforeUpdate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}
	return
}