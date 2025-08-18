package repository

import (
	"github.com/Rambudhi/payslip/internal/model"
	"gorm.io/gorm"
)

type AuthRepository interface {
	GetUserByUsername(username string) (*model.User, error)
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) GetUserByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
