package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Title     string `json:"title"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
