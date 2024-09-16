package routes

import (
	"net/http"
	"quiz-system/internal/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/admin/quizzes", controllers.CreateQuiz).Methods("POST")
	r.HandleFunc("/admin/quizzes/{id}", controllers.EditQuiz).Methods("PUT")
	r.HandleFunc("/admin/quizzes/{id}", controllers.DeleteQuiz).Methods("DELETE")

	r.HandleFunc("/quizzes", controllers.GetAllQuizzes).Methods("GET")
	r.HandleFunc("/quizzes/{id}/submit", controllers.SubmitQuiz).Methods("POST")

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods("GET")

	return r
}
