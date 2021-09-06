package entity

import (
	"time"

	"gorm.io/gorm"
)

type Interest struct {
	ID        uint64 `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `json:"name" gorm:"type:varchar(255);not null"`
}
