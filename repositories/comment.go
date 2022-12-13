package repositories

import (
	"fmt"
	"my-web-api/models"

	"github.com/jinzhu/gorm"
)

type CommentRepo interface {
	Create(comment *models.Comment) (*models.Comment, error)
	ReadAll() (*[]models.Comment, error)
	ReadById(id int) (*models.Comment, error)
	Update(id int, comment *models.Comment) (*models.Comment, error)
	Delete(id int) (map[string]interface{}, error)
}

type CommentRepoImpl struct {
	DB *gorm.DB
}

func CreateCommentRepo(DB *gorm.DB) CommentRepo {
	return &CommentRepoImpl{DB}
}

func (e *CommentRepoImpl) Create(comment *models.Comment) (*models.Comment, error) {
	result := e.DB.Model(&models.Comment{}).Create(comment)

	if result.Error != nil {
		return &models.Comment{}, fmt.Errorf("DB error : %v", result.Error)
	}

	return comment, nil
}

func (e *CommentRepoImpl) ReadAll() (*[]models.Comment, error) {
	comments := make([]models.Comment, 0)

	err := e.DB.Table("comments").Find(&comments).Error
	if err != nil {
		return nil, fmt.Errorf("DB error : %v", err)
	}

	return &comments, nil
}

func (e *CommentRepoImpl) ReadById(id int) (*models.Comment, error) {
	comment := &models.Comment{}

	errorRead := e.DB.Table("comments").Where("id = ?", id).First(comment).Error

	if errorRead != nil {
		return nil, errorRead
	}

	return comment, nil
}

func (e *CommentRepoImpl) Update(id int, comment *models.Comment) (*models.Comment, error) {
	updatedComment := &models.Comment{}
	result := e.DB.Model(updatedComment).Where("id = ?", id).Updates(models.Comment{
		PostId: comment.PostId, CommentId: comment.CommentId, Nickname: comment.Nickname, Content: comment.Content})

	if result.Error != nil {
		return nil, fmt.Errorf("DB error : %v", result.Error)
	}

	return updatedComment, nil
}

func (e *CommentRepoImpl) Delete(id int) (map[string]interface{}, error) {
	result := e.DB.Delete(&models.Comment{}, id)

	return map[string]interface{}{
		"rows_affected": result.RowsAffected,
	}, nil
}
