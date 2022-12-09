package implementations

import (
	"dot-crud-redis-go-api/models"
	"dot-crud-redis-go-api/repositories"
	"dot-crud-redis-go-api/usecases"
)

type CommentUsecaseImpl struct {
	commentRepo repositories.CommentRepo
}

func CreateCommentUsecase(commentRepo repositories.CommentRepo) usecases.CommentUsecase {
	return &CommentUsecaseImpl{commentRepo}
}

func (e *CommentUsecaseImpl) Create(comment *models.Comment) (*models.Comment, error) {
	return e.commentRepo.Create(comment)
}

func (e *CommentUsecaseImpl) ReadAll() (*[]models.Comment, error) {
	return e.commentRepo.ReadAll()
}

func (e *CommentUsecaseImpl) ReadById(id int) (*models.Comment, error) {
	return e.commentRepo.ReadById(id)
}

func (e *CommentUsecaseImpl) Update(id int, comment *models.Comment) (*models.Comment, error) {
	return e.commentRepo.Update(id, comment)
}

func (e *CommentUsecaseImpl) Delete(id int) (map[string]interface{}, error) {
	return e.commentRepo.Delete(id)
}
