package router

import (
	"log"
	"net/http"
	"todolist/config"
	"todolist/handler"
	"todolist/repository"
	"todolist/service"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	DB     *mongo.Collection = config.Connection()
	repo                     = repository.NewToDoList(DB)
	src                      = service.NewService(repo)
	handle                   = handler.NewHandler(src)
)

func API(r *gin.Engine) {
	api := r.Group("/api")

	{
		api.POST("/create", handle.CreateHandler)
	}

	config := config.GetConfiguration()
	log.Fatal(http.ListenAndServe(config.Port, r))
}
