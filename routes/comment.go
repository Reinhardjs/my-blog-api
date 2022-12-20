package routes

import (
	"my-web-api/configs"
	"my-web-api/controllers"

	"github.com/gorilla/mux"

	"my-web-api/repositories"
	"my-web-api/usecases"
)

func CommentRoute(router *mux.Router) {
	DB := configs.GetDB()

	commentRepository := repositories.CreateCommentRepo(DB)
	commentUsecase := usecases.CreateCommentUsecase(commentRepository)
	commentController := controllers.CreateCommentController(commentUsecase)

	router.Handle("/comments", commentController.GetComments()).Methods("GET")
	router.Handle("/comments/{commentId}", commentController.GetComment()).Methods("GET")
	// router.Handle("/comments", commentController.CreateComment()).Methods("POST")
	// router.Handle("/comments/{commentId}", commentController.UpdateComment()).Methods("PUT")
	// router.Handle("/comments/{commentId}", commentController.UpdateComment()).Methods("PATCH")
	// router.Handle("/comments/{commentId}", commentController.DeleteComment()).Methods("DELETE")
}
