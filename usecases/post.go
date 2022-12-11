package usecases

import (
	"dot-crud-redis-go-api/models"
	"dot-crud-redis-go-api/repositories"
)

type PostUsecase interface {
	Create(post *models.Post) (*models.Post, error)
	ReadAll() (*[]models.Post, error)
	ReadById(id int) (*models.Post, error)
	Update(id int, post *models.Post) (*models.Post, error)
	Delete(id int) (map[string]interface{}, error)
}

type PostUsecaseImpl struct {
	postRepo repositories.PostRepo
}

func CreatePostUsecase(postRepo repositories.PostRepo) PostUsecase {
	return &PostUsecaseImpl{postRepo}
}

func (e *PostUsecaseImpl) Create(post *models.Post) (*models.Post, error) {
	return e.postRepo.Create(post)
}

func (e *PostUsecaseImpl) ReadAll() (*[]models.Post, error) {
	return e.postRepo.ReadAll()
}

func (e *PostUsecaseImpl) ReadById(id int) (*models.Post, error) {
	return e.postRepo.ReadById(id)
}

func (e *PostUsecaseImpl) Update(id int, post *models.Post) (*models.Post, error) {
	return e.postRepo.Update(id, post)
}

func (e *PostUsecaseImpl) Delete(id int) (map[string]interface{}, error) {
	return e.postRepo.Delete(id)
}