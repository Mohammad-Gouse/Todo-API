package models

import "time"

type Todo struct {
	ID          string    `json:"id" bson:"_id,omitempty"`
	UserID      string    `json:"user_id" bson:"user_id"`
	Title       string    `json:"title" bson:"title"`
	Description string    `json:"description" bson:"description"`
	Status      string    `json:"status" bson:"status"`
	Created     time.Time `json:"created" bson:"created"`
	Updated     time.Time `json:"updated" bson:"updated"`
}
