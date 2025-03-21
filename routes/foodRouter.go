package routes

import (
	controller "golang-restuarant-management/controllers"

	"github.com/gin-gonic/gin"
)

func foodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controller.GetFoods())
	incomingRoutes.GET("/foods/:id", controller.GetFood())
	incomingRoutes.POST("/foods", controller.CreateFood())
	incomingRoutes.PATCH("/foods/:id", controller.UpdateFood())
}
