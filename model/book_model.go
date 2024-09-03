package model

import (
	"errors"

	uuid "github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	//gorm.Model

	ID         uuid.UUID `json:"id" gorm:"primarykey;type:uuid;default:gen_random_uuid()"` //search for uuid
	BookName   string    `json:"bookName"`                                                 //bookName
	BookAuthor string    `json:"bookAuthor"`
	BookQuant  int       `json:"bookQuant"`
}

type BookRepository interface {
	GetAllBooks() ([]Book, error)
	GetBookByID(id uuid.UUID) (Book, error)
	AddBook(book Book) error
	UpdateBook(id uuid.UUID, book Book) error
	DeleteBook(id uuid.UUID) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *DB) BookRepository {
	return &bookRepository{db: db.GormDB}
}

func (m *bookRepository) GetAllBooks() ([]Book, error) {
	var books []Book
	err := m.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (m *bookRepository) GetBookByID(id uuid.UUID) (Book, error) {
	var book Book
	err := m.db.First(&book, id).Error
	//first instead of find
	//if id is not there find will not return an error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Book{}, nil
		}
		return Book{}, err
	}
	return book, nil
}

func (m *bookRepository) AddBook(book Book) error {
	err := m.db.Create(&book).Error
	return err
}

func (m *bookRepository) UpdateBook(id uuid.UUID, book Book) error {
	res := m.db.Model(&Book{}).Where("id = ?", id).Updates(book)
	err := res.Error
	return err
}

func (m *bookRepository) DeleteBook(id uuid.UUID) error {
	err := m.db.Delete(&Book{}, id)
	return err.Error
}
