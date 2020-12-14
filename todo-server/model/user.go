package model

import (
	"beards.ly/todo/security"
	"gorm.io/gorm"
)

// User account in db
type User struct {
	gorm.Model
	Username   string `gorm:"size:320;not null;unique" json:"username"`
	Email      string `gorm:"size:320;not null;unique" json:"email"`
	Password   string `gorm:"size:100;not null" json:"password"`
	AvatarPath string `gorm:"size:2048;null" json:"avatar_path"`
}

// BeforeSave hook
func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
