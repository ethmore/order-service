package main

import (
	"order-service/dotEnv"
	"order-service/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{dotEnv.GoDotEnvVariable("BFFURL")}
	router.Use(cors.New(config))

	public := router.Group("/")
	routes.PrivateRoutes(public)

	router.Run(":3009")
}
