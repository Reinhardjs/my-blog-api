package models

import (
	"time"
)

type Post struct {
	ID        int64      `json:"id" gorm:"primary_key"`
	Url       string     `json:"url" gorm:"unique"`
	Nickname  string     `json:"nickname"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
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

	if post.Url == "" {
		return "Url should be on the payload", false
	}

	return "Payload is valid", true
}
