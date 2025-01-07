package entity

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID         uint           `gorm:"primaryKey"`
	Name       string         `gorm:"size:100;not null"`
	Price      float64        `gorm:"not null"`
	CategoryID uint           `gorm:"not null"`
	Category   Category       `gorm:"foreignKey:CategoryID"`
	CreatedAt  time.Time      `gorm:"autoCreateTime"`
	UpdatedAt  time.Time      `gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
