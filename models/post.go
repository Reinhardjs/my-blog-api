package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Url      string `gorm:"unique" json:"url"`
	Nickname string `json:"nickname"`
	Title    string `json:"title"`
	Content  string `json:"content"`
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
