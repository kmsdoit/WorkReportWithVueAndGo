package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(255)" json:"Email" binding:"required"`
	Password string `gorm:"type:text" json:"password"`
	Name     string `gorm:"type:text" json:"name"`
}
