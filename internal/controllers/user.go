package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"quiz-system/config"
	"quiz-system/internal/models"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllQuizzes(w http.ResponseWriter, r *http.Request) {
	collection := config.GetCollection("quizzes")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cur.Close(ctx)

	var quizzes []models.Quiz
	for cur.Next(ctx) {
		var quiz models.Quiz
		cur.Decode(&quiz)
		quizzes = append(quizzes, quiz)
	}

	json.NewEncoder(w).Encode(quizzes)
}

func SubmitQuiz(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	quizID, _ := primitive.ObjectIDFromHex(params["id"])

	var submittedAnswers map[string]string
	json.NewDecoder(r.Body).Decode(&submittedAnswers)

	collection := config.GetCollection("quizzes")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var quiz models.Quiz
	err := collection.FindOne(ctx, bson.M{"_id": quizID}).Decode(&quiz)
	if err == mongo.ErrNoDocuments {
		http.Error(w, "Viktorina topilmadi", http.StatusNotFound)
		return
	}

	var correctAnswers int
	for _, question := range quiz.Questions {
		if submittedAnswers[question.Question] == question.Answer {
			correctAnswers++
		}
	}

	score := float64(correctAnswers) / float64(len(quiz.Questions)) * 100
	result := map[string]interface{}{
		"score": score,
		"total": len(quiz.Questions),
	}

	json.NewEncoder(w).Encode(result)
}
