package handler

import (
	"encoding/json"
	"todolist/model"
	"todolist/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Handler struct {
	svc service.ToDoList
}

func NewHandler(svc service.ToDoList) *Handler {
	return &Handler{svc}
}

func (h *Handler) CreateHandler(c *gin.Context) {

	var input model.TodoListInput
	err := json.NewDecoder(c.Request.Body).Decode(&input)
	if err != nil {
		c.JSON(500, err)
		return
	}
	todo, err := h.svc.Create(input)
	result := bson.M{"_id": todo.ID, "task": input.Task, "status": false}
	if err != nil {
		c.JSON(400, err)
		return
	}
	json.NewEncoder(c.Writer).Encode(result)
}
