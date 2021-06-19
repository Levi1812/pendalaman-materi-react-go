package migration

import "time"

type User struct {
	ID            int `gorm:"primaryKey"`
	Name, Address string
	DateBirth     time.Time
	Email         string
	Password      string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Books         []Book `gorm:"foreignKey:UserID"`
}

type Book struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	Title     string
	Author    string
	Year      int
	CreatedAt time.Time
	UpdatedAt time.Time
}
