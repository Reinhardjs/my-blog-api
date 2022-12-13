package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"my-web-api/models"
	"my-web-api/responses"
	"my-web-api/usecases"
	"my-web-api/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type CommentController struct {
	commentsUsecase usecases.CommentUsecase
}

func CreateCommentController(commentsUsecase usecases.CommentUsecase) CommentController {
	return CommentController{commentsUsecase}
}

func (e *CommentController) GetComments() http.Handler {
	return RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		rw.Header().Add("Content-Type", "application/json")

		commentss, err := e.commentsUsecase.ReadAll()

		if err != nil {
			return err
		}

		response := responses.FineResponse{Status: http.StatusOK, Message: "success", Data: commentss}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}

func (e *CommentController) GetComment() http.Handler {
	return RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["commentId"])

		if err != nil {
			return utils.NewHTTPError(err, 400, "Invalid comments id")
		}

		rw.Header().Add("Content-Type", "application/json")

		comments, err := e.commentsUsecase.ReadById(int(id))

		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return utils.NewHTTPError(err, 404, "record not found")
			} else {
				return err
			}
		}

		response := responses.FineResponse{Status: http.StatusOK, Message: "success", Data: comments}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}

func (e *CommentController) CreateComment() http.Handler {
	return RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		rw.Header().Add("Content-Type", "application/json")

		comments := &models.Comment{}
		decodeError := json.NewDecoder(r.Body).Decode(comments)

		if decodeError != nil {
			return utils.NewHTTPError(nil, 400, "Invalid request body format")
		}

		result, err := e.commentsUsecase.Create(comments)

		if err != nil {
			return err
		}

		response := responses.FineResponse{Status: http.StatusCreated, Message: "success", Data: result}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}

func (e *CommentController) UpdateComment() http.Handler {
	return RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		rw.Header().Add("Content-Type", "application/json")

		params := mux.Vars(r)
		commentId, err := strconv.Atoi(params["commentId"])

		if err != nil {
			return utils.NewHTTPError(nil, 400, "Invalid comments id")
		}

		comments := &models.Comment{}
		decodeError := json.NewDecoder(r.Body).Decode(comments)
		if decodeError != nil {
			return utils.NewHTTPError(nil, 400, "Invalid request body format")
		}

		oldComment, oldCommentErr := e.commentsUsecase.ReadById(int(commentId))

		if oldCommentErr != nil {
			if errors.Is(oldCommentErr, gorm.ErrRecordNotFound) {
				return utils.NewHTTPError(oldCommentErr, 404, "record not found")
			} else {
				return oldCommentErr
			}
		}

		updatedComment, updateCommentErr := e.commentsUsecase.Update(commentId, comments)
		updatedComment.CreatedAt = oldComment.CreatedAt

		if updateCommentErr != nil {
			return updateCommentErr
		}

		if err != nil {
			return err
		}

		response := responses.FineResponse{Status: http.StatusOK, Message: "success", Data: updatedComment}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}

func (e *CommentController) DeleteComment() http.Handler {
	return RootHandler(func(rw http.ResponseWriter, r *http.Request) (err error) {
		_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		params := mux.Vars(r)
		id, err := strconv.Atoi(params["commentId"])

		if err != nil {
			return utils.NewHTTPError(err, 400, "Invalid comments id")
		}

		rw.Header().Add("Content-Type", "application/json")

		// Check for existing record
		_, existingCommentErr := e.commentsUsecase.ReadById(int(id))
		if existingCommentErr != nil {
			if errors.Is(existingCommentErr, gorm.ErrRecordNotFound) {
				return utils.NewHTTPError(existingCommentErr, 404, "record not found")
			} else {
				return existingCommentErr
			}
		}

		comments, err := e.commentsUsecase.Delete(id)

		if err != nil {
			return err
		}

		response := responses.FineResponse{Status: http.StatusOK, Message: "success", Data: comments}
		rw.WriteHeader(response.Status)
		json.NewEncoder(rw).Encode(response)
		return nil
	})
}
