package service

import (
	"errors"
	"golang2/model"

	"github.com/google/uuid"
)

type BookDTO struct {
	BookName   string `json:"bookName"`
	BookAuthor string `json:"bookAuthor"`
	BookQuant  int    `json:"bookQuant"`
}

type Bookservice interface {
	GetAllBooks() ([]model.Book, error)
	GetBookByID(id uuid.UUID) (model.Book, error)
	AddBook(book BookDTO) error
	UpdateBook(id uuid.UUID, book BookDTO) error
	DeleteBook(id uuid.UUID) error
}

type bookService struct {
	repo model.BookRepository
	//model.modelRepo
	//user.userrepo
	//repo model.modelrepo.book
	//user model.modelrepo.user
	//club all pacages of book and user in model
}

func NewBookService(db *model.DB) Bookservice {
	return &bookService{repo: model.NewBookRepository(db)}
}

func (s *bookService) GetAllBooks() ([]model.Book, error) {
	return s.repo.GetAllBooks()
}
func (s *bookService) GetBookByID(id uuid.UUID) (model.Book, error) {
	book, err := s.repo.GetBookByID(id)
	if err != nil {
		return book, errors.New("invalid ID")
	}
	if book.ID == uuid.Nil {
		return book, errors.New("no Book Found on the given ID")
	}
	return s.repo.GetBookByID(id)
}
func (s *bookService) AddBook(book BookDTO) error {
	return s.repo.AddBook(model.Book{
		BookName:   book.BookName,
		BookAuthor: book.BookAuthor,
		BookQuant:  book.BookQuant,
	})
}
func (s *bookService) UpdateBook(id uuid.UUID, book BookDTO) error {
	return s.repo.UpdateBook(id, model.Book{
		BookName:   book.BookName,
		BookAuthor: book.BookAuthor,
		BookQuant:  book.BookQuant,
	})
}
func (s *bookService) DeleteBook(id uuid.UUID) error {
	return s.repo.DeleteBook(id)
}
