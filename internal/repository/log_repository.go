package repository

import (
	"github.com/Rambudhi/payslip/internal/model"
	"gorm.io/gorm"
)

type LogRepository interface {
	Create(log *model.Log) error
}

type logRepository struct {
	db *gorm.DB
}

func NewLogRepository(db *gorm.DB) LogRepository {
	return &logRepository{db: db}
}

func (r *logRepository) Create(log *model.Log) error {
	return r.db.Create(log).Error
}
