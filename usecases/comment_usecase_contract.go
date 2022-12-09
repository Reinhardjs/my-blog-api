package usecases

import (
	"dot-crud-redis-go-api/models"
)

type CommentUsecase interface {
	Create(comment *models.Comment) (*models.Comment, error)
	ReadAll() (*[]models.Comment, error)
	ReadById(id int) (*models.Comment, error)
	Update(id int, comment *models.Comment) (*models.Comment, error)
	Delete(id int) (map[string]interface{}, error)
}
