package router

import (
	"gopkg.in/gin-gonic/gin.v1"
	"github.com/rakeshkumargupt/handler"
	"github.com/rakeshkumargupt/utils"
	"github.com/rakeshkumargupt/models"
)

func SetRoutes() *gin.Engine {
	router := gin.Default()

	// Endpoint call url : localhost:8080/sendwebpush
	router.POST("/sendwebpush", utils.Validator(models.UiWebpushModel{}), handler.SendWebpush)

	return router
}
