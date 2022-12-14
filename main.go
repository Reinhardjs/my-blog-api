package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"my-web-api/routes"
)

func main() {
	router := mux.NewRouter()

	// add routes
	routes.PostRoute(router)
	routes.CommentRoute(router)

	handler := cors.Default().Handler(router)

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", handler)
}
