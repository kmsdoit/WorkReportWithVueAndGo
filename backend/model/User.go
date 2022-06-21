package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	JoinCompanyYear *time.Time `json:"JoinCompanyYear"`
	Image           string     `gorm:"type:varchar(255)" json:"image" binding:"required"`
	EmpId           string     `gorm:"type:varchar(255)" json:"EmpId" binding:"required"`
	Email           string     `gorm:"type:varchar(255)" json:"Email" binding:"required"`
	Password        string     `gorm:"type:text" json:"password"`
	Name            string     `gorm:"type:text" json:"name"`
}
