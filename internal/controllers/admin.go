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
)

func CreateQuiz(w http.ResponseWriter, r *http.Request) {
	var quiz models.Quiz
	json.NewDecoder(r.Body).Decode(&quiz)

	collection := config.GetCollection("quizzes")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, quiz)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func EditQuiz(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	quizID, _ := primitive.ObjectIDFromHex(params["id"])

	var quiz models.Quiz
	json.NewDecoder(r.Body).Decode(&quiz)

	collection := config.GetCollection("quizzes")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{
		"$set": quiz,
	}

	_, err := collection.UpdateOne(ctx, bson.M{"_id": quizID}, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteQuiz(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	quizID, _ := primitive.ObjectIDFromHex(params["id"])

	collection := config.GetCollection("quizzes")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.DeleteOne(ctx, bson.M{"_id": quizID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
