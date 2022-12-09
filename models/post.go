package models

import (
	"time"
)

type Post struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `json:"nickname"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (post *Post) Validate() (string, bool) {

	if post.Title == "" {
		return "Title should be on the payload", false
	}

	if post.Nickname == "" {
		return "Nickname should be on the payload", false
	}

	if post.Content == "" {
		return "Content should be on the payload", false
	}

	return "Payload is valid", true
}
