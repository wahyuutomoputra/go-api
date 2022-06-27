package book

import "gorm.io/gorm"

type BookRepository interface {
	FindAll() ([]Book, error)
}

type bookRepository struct {
	db *gorm.DB
}

func NewRepositoryBook(db *gorm.DB) *bookRepository {
	return &bookRepository{db}
}

func (s *bookRepository) FindAll() ([]Book, error) {
	var book []Book
	err := s.db.Find(&book).Error

	return book, err
}
