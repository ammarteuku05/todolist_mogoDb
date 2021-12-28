package repository

import (
	"context"
	"todolist/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ToDoList interface {
	Create(model.ToDoList) (model.ToDoList, error)
}

type repository struct {
	db *mongo.Collection
}

func NewToDoList(db *mongo.Collection) *repository {
	return &repository{db}
}

func (r *repository) Create(mod model.ToDoList) (model.ToDoList, error) {
	res, err := r.db.InsertOne(context.Background(), mod)

	todo := model.ToDoList{
		ID: res.InsertedID.(primitive.ObjectID),
	}

	if err != nil {
		return todo, err
	}
	return todo, nil

}
