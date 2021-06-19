package book

import "gorm.io/gorm"

type Repository interface {
	GetAll() ([]Book, error)
	FindByID(ID string) (Book, error)
	FindByUserID(userID string) ([]Book, error)
	Create(book Book) (Book, error)
	Update(ID string, dataUpdate map[string]interface{}) (Book, error)
	Delete(ID string) (string, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Book, error) {
	var books []Book

	if err := r.db.Find(&books).Error; err != nil {
		return books, err
	}

	return books, nil
}

func (r *repository) FindByID(ID string) (Book, error) {
	var book Book

	if err := r.db.Where("id = ?", ID).Find(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}
func (r *repository) FindByUserID(userID string) ([]Book, error) {
	var books []Book

	if err := r.db.Where("user_id = ?", userID).Find(&books).Error; err != nil {
		return books, err
	}

	return books, nil
}
func (r *repository) Create(book Book) (Book, error) {
	if err := r.db.Create(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) Update(ID string, dataUpdate map[string]interface{}) (Book, error) {
	var book Book

	if err := r.db.Model(&book).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return book, err
	}

	if err := r.db.Where("id = ?", ID).Find(&book).Error; err != nil {
		return book, err
	}

	return book, nil
}

func (r *repository) Delete(ID string) (string, error) {
	var book Book

	if err := r.db.Where("id = ?", ID).Delete(&book).Error; err != nil {
		return "error", err
	}

	return "success", nil
}
