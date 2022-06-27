package main

import (
	"go-api/src/config"
	"go-api/src/modules/book"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	db, err := config.ConnectDb()
	if err != nil {
		panic("Failed to connect to database!")
	}

	bookRepository := book.NewRepositoryBook(db)
	bookService := book.NewServiceBook(bookRepository)
	bookController := book.NewBookController(bookService)

	v1 := router.Group("/v1")
	{
		v1.GET("/status", bookController.Status)
		v1.GET("/findAll", bookController.FindAll)
	}

	router.Run(":8081")
}
