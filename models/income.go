package models

import (
	"time"

	"gorm.io/gorm"
)

type Income struct {
	ID          uint     `json:"id" gorm:"primaryKey"`
	Amount      float64  `json:"amount"`
	Description string   `json:"description"`
	CategoryID  uint     `json:category`
	Category    Category `json:category gorm:foreignKey:CompanyRefer`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
