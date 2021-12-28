package model

type TodoListInput struct {
	Task string `json:"task" bson:"task"`
}
