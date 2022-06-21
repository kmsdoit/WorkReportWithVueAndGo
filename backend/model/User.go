package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	EmpId    string `gorm:"type:varchar(255)" json:"EmpId" binding:"required"`
	Email    string `gorm:"type:varchar(255)" json:"Email" binding:"required"`
	Password string `gorm:"type:text" json:"password"`
	Name     string `gorm:"type:text" json:"name"`
}
