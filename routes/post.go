package routes

import (
	"my-web-api/configs"
	"my-web-api/controllers"

	"github.com/gorilla/mux"

	"my-web-api/repositories"
	"my-web-api/usecases"
)

func PostRoute(router *mux.Router) {
	DB := configs.GetDB()
	RedisClient := configs.GetRedis()

	postRepository := repositories.CreatePostRepo(DB, RedisClient)
	postUsecase := usecases.CreatePostUsecase(postRepository)
	postController := controllers.CreatePostController(postUsecase)

	router.Handle("/posts", postController.GetPosts()).Methods("GET")
	router.Handle("/posts/{postUrl}", postController.GetPost()).Methods("GET")
	// router.Handle("/posts", postController.CreatePost()).Methods("POST")
	// router.Handle("/posts/{postUrl}", postController.UpdatePost()).Methods("PUT")
	// router.Handle("/posts/{postUrl}", postController.UpdatePost()).Methods("PATCH")
	// router.Handle("/posts/{postUrl}", postController.DeletePost()).Methods("DELETE")
	router.Handle("/{postTag}", postController.GetPostsByTag()).Methods("GET")
}
