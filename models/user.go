package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"

	"rakaminbtpn/helpers"
)

type User struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	Username  string     `gorm:"not null" json:"username" form:"username" valid:"required~Your username is required"`
	Email     string     `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password  string     `gorm:"not null" json:"-" form:"password" valid:"required~Your password is required,minstringlength(6)~Password has to have a minimum length of 6 characters"`
	Photos    []Photo    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"-"`
	CreatedAt *time.Time `json:"-,omitempty"`
	UpdatedAt *time.Time `json:"-,omitempty"`
}

func (u *User) Before_create(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.Hash_pass(u.Password)

	err = nil
	return
}

func (u *User) Before_update(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	u.Password = helpers.Hash_pass(u.Password)
	err = nil
	return
}
