package repositories

import (
	"dot-crud-redis-go-api/models"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/garyburd/redigo/redis"
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
	DB          *gorm.DB
	RedisClient redis.Conn
}

func CreateCommentRepo(DB *gorm.DB, RedisClient redis.Conn) CommentRepo {
	return &CommentRepoImpl{DB, RedisClient}
}

func (e *CommentRepoImpl) Create(comment *models.Comment) (*models.Comment, error) {
	result := e.DB.Model(&models.Comment{}).Create(comment)

	if result.Error != nil {
		return &models.Comment{}, fmt.Errorf("DB error : %v", result.Error)
	}

	_, redisDeleteAllErr := e.RedisClient.Do("DEL", "comment:all")

	if redisDeleteAllErr != nil {
		// Failed deleting data (comment:all) from redis
		return nil, redisDeleteAllErr
	}

	return comment, nil
}

func (e *CommentRepoImpl) ReadAll() (*[]models.Comment, error) {
	comments := make([]models.Comment, 0)

	// Get JSON blob from Redis
	redisResult, err := e.RedisClient.Do("GET", "comment:all")

	if err != nil {
		// Failed getting data from redis
		return nil, err
	}

	if redisResult == nil {

		err := e.DB.Table("comments").Find(&comments).Error
		if err != nil {
			return nil, fmt.Errorf("DB error : %v", err)
		}

		commentJSON, err := json.Marshal(comments)
		if err != nil {
			return nil, err
		}

		// Save JSON blob to Redis
		_, saveRedisError := e.RedisClient.Do("SET", "comment:all", commentJSON)

		if saveRedisError != nil {
			// Failed saving data to redis
			return nil, saveRedisError
		}
	} else {
		json.Unmarshal(redisResult.([]byte), &comments)
	}

	return &comments, nil
}

func (e *CommentRepoImpl) ReadById(id int) (*models.Comment, error) {
	comment := &models.Comment{}

	// Get JSON blob from Redis
	redisResult, err := e.RedisClient.Do("GET", "comment:"+strconv.Itoa(id))

	if err != nil {
		// Failed getting data from redis
		return nil, err
	}

	if redisResult == nil {

		errorRead := e.DB.Table("comments").Where("id = ?", id).First(comment).Error

		if errorRead != nil {
			return nil, errorRead
		}

		commentJSON, err := json.Marshal(comment)
		if err != nil {
			return nil, err
		}

		// Save JSON blob to Redis
		_, saveRedisError := e.RedisClient.Do("SET", "comment:"+strconv.Itoa(id), commentJSON)

		if saveRedisError != nil {
			// Failed saving data to redis
			return nil, saveRedisError
		}
	} else {
		json.Unmarshal(redisResult.([]byte), &comment)
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

	// Delete JSON blob from Redis
	_, redisDeleteErr := e.RedisClient.Do("DEL", "comment:"+strconv.Itoa(id))
	_, redisDeleteAllErr := e.RedisClient.Do("DEL", "comment:all")

	if redisDeleteErr != nil {
		// Failed deleting data from redis
		return nil, redisDeleteErr
	}

	if redisDeleteAllErr != nil {
		// Failed deleting data (comment:all) from redis
		return nil, redisDeleteAllErr
	}

	return updatedComment, nil
}

func (e *CommentRepoImpl) Delete(id int) (map[string]interface{}, error) {
	result := e.DB.Delete(&models.Comment{}, id)

	// Delete JSON blob from Redis
	_, redisDeleteErr := e.RedisClient.Do("DEL", "comment:"+strconv.Itoa(id))
	_, redisDeleteAllErr := e.RedisClient.Do("DEL", "comment:all")

	if redisDeleteErr != nil {
		// Failed deleting data from redis
		return nil, redisDeleteErr
	}

	if redisDeleteAllErr != nil {
		// Failed deleting data (comment:all) from redis
		return nil, redisDeleteAllErr
	}

	return map[string]interface{}{
		"rows_affected": result.RowsAffected,
	}, nil
}
