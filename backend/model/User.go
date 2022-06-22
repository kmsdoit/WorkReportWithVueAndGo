package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	JoinCompanyYear *time.Time `json:"JoinCompanyYear"`                                   //입사일
	EmpId           string     `gorm:"type:varchar(255)" json:"EmpId" binding:"required"` //사원 아이디
	Email           string     `gorm:"type:varchar(255)" json:"Email" binding:"required"` // 사원 메일
	Password        string     `gorm:"type:text" json:"password"`                         //비밀번호
	Name            string     `gorm:"type:text" json:"name"`                             // 이름
}
