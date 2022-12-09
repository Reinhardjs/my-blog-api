package models

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	PostId    int    `json:"post_id"`
	CommentId int    `json:"comment_id"`
	Nickname  string `json:"nickname"`
	Content   string `json:"content"`
}
