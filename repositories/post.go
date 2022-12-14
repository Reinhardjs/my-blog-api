package repositories

import (
	"fmt"
	"my-web-api/models"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
)

type PostRepo interface {
	Create(post *models.Post) (*models.Post, error)
	ReadAll() (*[]models.Post, error)
	ReadById(id int) (*models.Post, error)
	Update(id int, post *models.Post) (*models.Post, error)
	Delete(id int) (map[string]interface{}, error)
}

type PostRepoImpl struct {
	DB          *gorm.DB
	RedisClient redis.Conn
}

func CreatePostRepo(DB *gorm.DB, RedisClient redis.Conn) PostRepo {
	return &PostRepoImpl{DB, RedisClient}
}

func (e *PostRepoImpl) Create(post *models.Post) (*models.Post, error) {
	result := e.DB.Model(&models.Post{}).Create(post)

	if result.Error != nil {
		return &models.Post{}, fmt.Errorf("DB error : %v", result.Error)
	}

	return post, nil
}

func (e *PostRepoImpl) ReadAll() (*[]models.Post, error) {
	posts := make([]models.Post, 0)

	err := e.DB.Table("posts").Find(&posts).Error
	if err != nil {
		return nil, fmt.Errorf("DB error : %v", err)
	}

	return &posts, nil
}

func (e *PostRepoImpl) ReadById(id int) (*models.Post, error) {
	post := &models.Post{}

	errorRead := e.DB.Table("posts").Where("id = ?", id).First(post).Error

	if errorRead != nil {
		return nil, errorRead
	}

	return post, nil
}

func (e *PostRepoImpl) Update(id int, post *models.Post) (*models.Post, error) {
	updatedPost := &models.Post{}
	result := e.DB.Model(updatedPost).Where("id = ?", id).Updates(models.Post{Nickname: post.Nickname, Title: post.Title, Content: post.Content})

	if result.Error != nil {
		return nil, fmt.Errorf("DB error : %v", result.Error)
	}

	return updatedPost, nil
}

func (e *PostRepoImpl) Delete(id int) (map[string]interface{}, error) {
	result := e.DB.Delete(&models.Post{}, id)

	return map[string]interface{}{
		"rows_affected": result.RowsAffected,
	}, nil
}
