package routes

import (
	"go-api/src/modules/book"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func bookRouter(r *gin.RouterGroup, db *gorm.DB) {
	bookRepository := book.NewRepositoryBook(db)
	bookService := book.NewServiceBook(bookRepository)
	bookController := book.NewBookController(bookService)

	router := r.Group("/book")
	{
		router.GET("/findAll", bookController.FindAll)
		router.POST("/create", bookController.Create)
	}
}

func AddRoutes(r *gin.RouterGroup, db *gorm.DB) {
	bookRouter(r, db)
}
