package book

type BookService interface {
	FindAll() ([]Book, error)
	Insert(book Book) (Book, error)
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

func (s *bookService) Insert(book Book) (Book, error) {
	return s.bookRepository.Insert(book)
}
