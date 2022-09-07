package routes

import (
	"order/controllers"

	"github.com/gin-gonic/gin"
)

func PrivateRoutes(g *gin.RouterGroup) {
	g.POST("/createOrder", controllers.CreateOrder())
}
