package routes

import (
	"order/controllers"

	"github.com/gin-gonic/gin"
)

func PrivateRoutes(g *gin.RouterGroup) {
	g.GET("/test", controllers.Test())

	g.POST("/createOrder", controllers.CreateOrder())
}
