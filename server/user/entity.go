package user

import (
	"backendGolang/book"
	"time"
)

type User struct {
	ID            int `gorm:"primaryKey"`
	Name, Address string
	DateBirth     time.Time
	Email         string
	Password      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Books         []book.Book `gorm:"foreignKey:UserID"`
}
