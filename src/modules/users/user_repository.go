package users

import "gorm.io/gorm"

type UserRepository interface {
	FindAll() ([]Users, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewRepositoryUsers(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r userRepository) FindAll() ([]Users, error) {
	var users []Users
	err := r.db.Find(&users).Error

	return users, err
}
