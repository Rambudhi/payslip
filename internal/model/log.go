package model

import (
	"time"
)

type Log struct {
	ID          uint   `gorm:"primaryKey"`
	TableName   string `gorm:"not null"`
	RecordID    uint   `gorm:"not null"`
	Action      string `gorm:"type:varchar(20);not null;check:action IN ('create','update','delete')"`
	OldData     string `gorm:"type:jsonb"`
	NewData     string `gorm:"type:jsonb"`
	PerformedBy *uint
	IPAddress   string
	RequestID   string
	CreatedAt   time.Time
}
