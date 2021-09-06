package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID                        uint64 `gorm:"primaryKey"`
	CreatedAt                 time.Time
	UpdatedAt                 time.Time
	DeletedAt                 gorm.DeletedAt `gorm:"index"`
	Name                      string         `json:"name" gorm:"type:varchar(100);not null"`
	Username                  string         `json:"username" gorm:"type:varchar(100);not null;uniqueIndex"`
	Email                     string         `json:"email" gorm:"type:varchar(100);not null;uniqueIndex"`
	Phone                     string         `json:"phone" gorm:"type:varchar(100);not null;uniqueIndex"`
	PhoneCode                 string         `json:"phone_code" gorm:"type:varchar(5);not null"`
	Gender                    string         `json:"gender" gorm:"type:varchar(100)"`
	Bio                       string         `json:"bio"`
	Website                   string         `json:"website" gorm:"type:varchar(100)"`
	Password                  string         `json:"-" gorm:"->;<-;not null"`
	IsAdmin                   bool           `json:"is_admin" gorm:"default:false"`
	IsAccountSuggestions      bool           `json:"is_account_suggestions" gorm:"default:true"`
	IsPrivate                 bool           `json:"is_private" gorm:"default:false"`
	IsTwoFactorAuthentication bool           `json:"is_two_factor_authentication" gorm:"default:false"`
}
