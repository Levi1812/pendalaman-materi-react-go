package book

import "time"

type Book struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Year      int       `json:"year"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
