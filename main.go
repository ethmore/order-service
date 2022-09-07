package main

import (
	"order/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3001"}
	router.Use(cors.New(config))

	public := router.Group("/")
	routes.PrivateRoutes(public)

	router.Run(":3009")
}
