package models

import "time"

type Comment struct {
	ID        int       `gorm:"primary_key;auto_increment" json:"id"`
	PostId    int       `json:"post_id"`
	CommentId int       `json:"comment_id"`
	Nickname  string    `json:"nickname"`
	Content   string    `json:"content"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
