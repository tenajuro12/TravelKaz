package models

import (
	"time"
)

// Blog Model
type Blog struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	Likes     int       `json:"likes"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Comments  []Comment `gorm:"foreignKey:BlogID" json:"comments"`
}

// Comment Model
type Comment struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	BlogID    uint      `json:"blog_id"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// Like Model
type Like struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	BlogID uint `json:"blog_id"`
	UserID uint `json:"user_id"`
}
