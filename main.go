package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"dot-crud-redis-go-api/routes"
)

func main() {
	router := mux.NewRouter()

	// add routes
	routes.PostRoute(router)
	routes.CommentRoute(router)

	fmt.Println("starting web server at http://localhost:8080")
	http.ListenAndServe(":8080", router)
}
