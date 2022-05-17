package model

import (
	"golang_api/helper"

	"gorm.io/gorm"
)

type User struct {
	ID       int
	Password string `json:"password" binding:"required"`
	Username string `json:"username" binding:"required"`
}

const TableUser = "user"

func (User) TableName() string {
	return TableUser
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	hashPass, _ := helper.GeneratePassword(u.Password)
	u.Password = hashPass
	return
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	hashPass, _ := helper.GeneratePassword(u.Password)
	u.Password = hashPass
	return
}
