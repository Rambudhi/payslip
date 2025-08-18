package repository

import (
	"github.com/Rambudhi/payslip/internal/model"

	"gorm.io/gorm"
)

type PayrollPeriodRepository interface {
	Create(period *model.PayrollPeriod) error
}

type payrollPeriodRepository struct {
	db *gorm.DB
}

func NewPayrollPeriodRepository(db *gorm.DB) PayrollPeriodRepository {
	return &payrollPeriodRepository{db: db}
}

func (r *payrollPeriodRepository) Create(period *model.PayrollPeriod) error {
	return r.db.Create(period).Error
}
