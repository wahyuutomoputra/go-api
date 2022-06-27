package book

import "gorm.io/gorm"

type BookRepository interface {
	FindAll() ([]Book, error)
	Insert(book Book) (Book, error)
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

func (s *bookRepository) Insert(book Book) (Book, error) {
	err := s.db.Debug().Create(&book).Error
	if err != nil {
		panic(err.Error())
	}
	return book, err
}
