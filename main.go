package main

import (
	"go-api/src/config"
	"go-api/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db, err := config.ConnectDb()
	if err != nil {
		panic("Failed to connect to database!")
	}

	r := router.Group("/v1")
	{
		routes.AddRoutes(r, db)
	}

	router.Run(":8081")
}
