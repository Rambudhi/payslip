package model

import (
	"time"
)

type User struct {
	ID        uint    `gorm:"primaryKey"`
	Username  string  `gorm:"unique;not null"`
	Password  string  `gorm:"not null"`
	Role      string  `gorm:"type:varchar(20);not null;check:role IN ('admin','employee')"`
	Salary    float64 `gorm:"not null;default:0"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
