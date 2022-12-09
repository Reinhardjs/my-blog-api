package routes

import (
	"dot-crud-redis-go-api/configs"
	"dot-crud-redis-go-api/controllers"

	"github.com/gorilla/mux"

	repositories "dot-crud-redis-go-api/repositories/implementations"
	usecases "dot-crud-redis-go-api/usecases/implementations"
)

func CommentRoute(router *mux.Router) {
	DB := configs.GetDB()
	RedisClient := configs.GetRedis()

	commentRepository := repositories.CreateCommentRepo(DB, RedisClient)
	commentUsecase := usecases.CreateCommentUsecase(commentRepository)
	commentController := controllers.CreateCommentController(commentUsecase)

	router.Handle("/comments", commentController.GetComments()).Methods("GET")
	router.Handle("/comments/{commentId}", commentController.GetComment()).Methods("GET")
	router.Handle("/comments", commentController.CreateComment()).Methods("POST")
	router.Handle("/comments/{commentId}", commentController.UpdateComment()).Methods("PUT")
	router.Handle("/comments/{commentId}", commentController.UpdateComment()).Methods("PATCH")
	router.Handle("/comments/{commentId}", commentController.DeleteComment()).Methods("DELETE")
}
