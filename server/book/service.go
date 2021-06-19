package book

import (
	"fmt"
	"time"
)

type Service interface {
	GetAllBookByUser(userID string) ([]Book, error)
	Createbook(userID int, input CreateBook) (Book, error)
	GetBookByID(bookID string) (Book, error)
	UpdateBook(bookID string, input UpdateBook) (Book, error)
	DeleteBook(bookID string) (string, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllBookByUser(userID string) ([]Book, error) {
	books, err := s.repository.FindByUserID(userID)

	if err != nil {
		return books, err
	}

	return books, nil
}
func (s *service) Createbook(userID int, input CreateBook) (Book, error) {

	var newBook = Book{
		UserID:    userID,
		Title:     input.Title,
		Author:    input.Author,
		Year:      input.Year,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	book, err := s.repository.Create(newBook)

	if err != nil {
		return book, err
	}

	return book, nil
}

func (s *service) GetBookByID(bookID string) (Book, error) {
	book, err := s.repository.FindByID(bookID)

	if err != nil {
		return book, err
	}

	return book, nil
}

func (s *service) UpdateBook(bookID string, input UpdateBook) (Book, error) {
	var dataUpdate = map[string]interface{}{}

	if input.Title != "" || len(input.Title) != 0 {
		dataUpdate["title"] = input.Title
	}

	if input.Author != "" || len(input.Author) != 0 {
		dataUpdate["author"] = input.Author
	}

	if input.Year != 0 {
		dataUpdate["year"] = input.Year
	}

	dataUpdate["updated_at"] = time.Now()

	book, err := s.repository.Update(bookID, dataUpdate)

	if err != nil {
		return book, err
	}

	return book, nil

}
func (s *service) DeleteBook(bookID string) (string, error) {
	msg, err := s.repository.Delete(bookID)

	if err != nil || msg == "error" {
		return msg, err
	}

	message := fmt.Sprintf("book id %s success deleted", bookID)

	return message, nil
}
