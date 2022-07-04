package routes

import (
	"go-api/src/modules/book"
	"go-api/src/modules/users"

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
		router.GET("/status", bookController.Status)
	}
}

func usersRouter(r *gin.RouterGroup, db *gorm.DB) {
	usersRepository := users.NewRepositoryUsers(db)
	usersService := users.NewServiceUser(usersRepository)
	usersControllers := users.NewUserController(usersService)

	router := r.Group("/users")
	{
		router.GET("/findAll", usersControllers.FindAll)
	}
}

func AddRoutes(r *gin.RouterGroup, db *gorm.DB) {
	bookRouter(r, db)
	usersRouter(r, db)
}
