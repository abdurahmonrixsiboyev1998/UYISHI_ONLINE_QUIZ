package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Question struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Question string             `bson:"question" json:"question"`
	Options  []string           `bson:"options" json:"options"`
	Answer   string             `bson:"answer" json:"answer"`
}

type Quiz struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Questions   []Question         `bson:"questions" json:"questions"`
}
