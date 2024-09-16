package main

import (
	"log"
	"net/http"
	"quiz-system/config"
	"quiz-system/internal/routes"
)

func main() {
	config.ConnectDB()

	r := routes.SetupRoutes()

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
