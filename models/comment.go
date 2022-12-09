package models

import (
	"time"
)

type Comment struct {
	ID        int64      `json:"id" gorm:"primary_key"`
	PostId    int        `json:"post_id"`
	CommentId int        `json:"comment_id"`
	Nickname  string     `json:"nickname"`
	Content   string     `json:"content"`
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt" sql:"index"`
}
