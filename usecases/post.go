package usecases

import (
	"my-web-api/models"
	"my-web-api/repositories"
)

type PostUsecase interface {
	Create(post *models.Post) (*models.Post, error)
	ReadAll() (*[]models.Post, error)
	ReadByUrl(url string) (*models.Post, error)
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

func (e *PostUsecaseImpl) ReadByUrl(url string) (*models.Post, error) {
	return e.postRepo.ReadByUrl(url)
}

func (e *PostUsecaseImpl) Update(id int, post *models.Post) (*models.Post, error) {
	return e.postRepo.Update(id, post)
}

func (e *PostUsecaseImpl) Delete(id int) (map[string]interface{}, error) {
	return e.postRepo.Delete(id)
}
