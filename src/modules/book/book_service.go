package book

type BookService interface {
	FindAll() ([]Book, error)
}

type bookService struct {
	bookRepository BookRepository
}

func NewServiceBook(bookRepository BookRepository) *bookService {
	return &bookService{bookRepository}
}

func (s *bookService) FindAll() ([]Book, error) {
	return s.bookRepository.FindAll()
}
