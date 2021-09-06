package entity

import "gorm.io/gorm"

type Token struct {
	gorm.Model
	Token    string `json:"token" gorm:"type:varchar(100)"`
	Agent    string `json:"agent"`
	ClientIP string `json:"client_ip"`
	UserId   uint64 `json:"user_id"`
}
