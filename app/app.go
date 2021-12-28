package app

import (
	"todolist/router"

	"github.com/gin-gonic/gin"
)

var (
	rout = gin.Default()
)

func StartApplication() {
	router.API(rout)

}
