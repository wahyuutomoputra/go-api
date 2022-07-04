package main

import (
	"go-api/src/config"
	"go-api/src/helper"
	"go-api/src/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, helper.ApiResponse(http.StatusNotFound, "Not found", nil))
	})

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
